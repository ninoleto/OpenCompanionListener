'use strict'
const test = require('node:test')
const assert = require('node:assert/strict')
const http = require('node:http')
const {request, makeHeaders} = require('./transport.cjs')
const key = 'ab'.repeat(32) // Deliberately synthetic, never a deployment key.
async function fixture(t, handler) {
  const server = http.createServer(handler)
  await new Promise(resolve=>server.listen(0,'127.0.0.1',resolve))
  t.after(()=>{server.closeAllConnections();server.close()})
  return {host:'127.0.0.1',port:server.address().port}
}
test('signed HTTP transport preserves method, path and body',async t=>{
  const command={type:'scroll',axis:'vertical',steps:-5}
  const config=await fixture(t,(req,res)=>{
    const chunks=[]
    req.on('data',c=>chunks.push(c))
    req.on('end',()=>{
      const body=Buffer.concat(chunks)
      const signed=makeHeaders(key,req.method,req.url,body,req.headers['x-ocl-timestamp'],req.headers['x-ocl-nonce'])
      assert.equal(req.headers['x-ocl-signature'],signed['X-OCL-Signature'])
      assert.deepEqual(JSON.parse(body),command)
      res.end('{"ok":true}')
    })
  })
  assert.deepEqual(await request(config,key,'POST','/v1/command',command),{ok:true})
})
test('truncated responses reject instead of remaining pending',async t=>{
  const config=await fixture(t,(_req,res)=>{
    res.writeHead(200,{'Content-Length':'100'})
    res.write('{')
    setTimeout(()=>res.destroy(),20)
  })
  let timer
  try {
    const result=await Promise.race([
      request(config,key,'GET','/v1/status').then(()=> 'resolved',e=>e.message),
      new Promise(resolve=>{timer=setTimeout(()=>resolve('pending'),4500)})
    ])
    assert.match(result,/interrupted|aborted|closed/i)
  } finally {clearTimeout(timer)}
})
test('oversized responses reject rather than accepting a truncated JSON prefix',async t=>{
  const config=await fixture(t,(_req,res)=>{
    res.write('{"ok":true}')
    setTimeout(()=>res.end(' '.repeat(9000)),20)
  })
  await assert.rejects(request(config,key,'GET','/v1/status'),/too large/i)
})
test('malformed JSON and non-2xx responses report errors',async t=>{
  for(const [code,body,pattern] of [[200,'invalid',/invalid JSON/],[401,'authentication failed',/HTTP 401/]]) {
    const config=await fixture(t,(_req,res)=>{res.writeHead(code);res.end(body)})
    await assert.rejects(request(config,key,'GET','/v1/status'),pattern)
  }
})
test('idle server times out and releases request',async t=>{
  const config=await fixture(t,()=>{})
  await assert.rejects(request(config,key,'GET','/v1/status'),/timed out/i)
})
test('response deadline also bounds a continuously streaming server',async t=>{
  const intervals=new Set()
  t.after(()=>{for(const id of intervals)clearInterval(id)})
  const config=await fixture(t,(_req,res)=>{
    res.write('{')
    const interval=setInterval(()=>res.write(' '),50);intervals.add(interval)
    res.on('close',()=>{clearInterval(interval);intervals.delete(interval)})
  })
  let timer
  try {
    const result=await Promise.race([
      request(config,key,'GET','/v1/status').then(()=> 'resolved',e=>e.message),
      new Promise(resolve=>{timer=setTimeout(()=>resolve('pending'),4500)})
    ])
    assert.match(result,/timed out/i)
  } finally {clearTimeout(timer)}
})
