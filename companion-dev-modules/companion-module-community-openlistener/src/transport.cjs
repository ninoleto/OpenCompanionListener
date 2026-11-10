'use strict'
const crypto = require('node:crypto')
const http = require('node:http')

const agent = new http.Agent({ keepAlive: true, maxSockets: 4, timeout: 10000 })
function makeHeaders(keyHex, method, path, body, now = Math.floor(Date.now() / 1000), nonce = crypto.randomBytes(16).toString('hex')) {
  if (!/^[\da-f]{64}$/i.test(keyHex)) throw new Error('Pairing key must contain 64 hex characters')
  const digest = crypto.createHash('sha256').update(body).digest('hex')
  const payload = [method, path, String(now), nonce, digest].join('\n')
  const signature = crypto.createHmac('sha256', Buffer.from(keyHex, 'hex')).update(payload).digest('hex')
  return { 'X-OCL-Timestamp': String(now), 'X-OCL-Nonce': nonce, 'X-OCL-Signature': signature }
}
async function request(config, keyHex, method, path, command = null) {
  const body = command === null ? Buffer.alloc(0) : Buffer.from(JSON.stringify(command), 'utf8')
  const headers = makeHeaders(keyHex, method, path, body)
  headers['Content-Type'] = 'application/json'
  headers['Content-Length'] = body.length
  const hostname = String(config.host || '').trim()
  const port = Number(config.port || 17992)
  if (!/^(?:\d{1,3}\.){3}\d{1,3}$/.test(hostname) || port < 1 || port > 65535) throw new Error('Invalid listener IP / port')
  return new Promise((resolve, reject) => {
    const req = http.request({ hostname, port, method, path, headers, agent, timeout: 3500 }, res => {
      const chunks = []; let size = 0
      res.on('data', chunk => {size += chunk.length; if (size <= 8192) chunks.push(chunk) })
      res.on('end', () => {
        const raw = Buffer.concat(chunks).toString('utf8')
        if (res.statusCode < 200 || res.statusCode >= 300) return reject(new Error(`Listener HTTP ${res.statusCode}: ${raw.slice(0,200)}`))
        try { resolve(JSON.parse(raw)) } catch { reject(new Error('Listener sent invalid JSON')) }
      })
    })
    req.on('timeout', () => req.destroy(new Error('Listener timed out')))
    req.on('error', reject)
    req.end(body)
  })
}
module.exports = { makeHeaders, request }
