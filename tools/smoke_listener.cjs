'use strict'
// Limited Windows binary smoke check. No desktop input, app launching, browser
// opening, startup writes or live Companion use. Fixtures stay under .cache/.
const assert = require('node:assert/strict')
const crypto = require('node:crypto')
const fs = require('node:fs')
const http = require('node:http')
const net = require('node:net')
const path = require('node:path')
const {spawn} = require('node:child_process')
const {setTimeout: delay} = require('node:timers/promises')
const {makeHeaders, request} = require('../companion-dev-modules/companion-module-community-openlistener/src/transport.cjs')

async function freePort(port) {
  const server = net.createServer()
  await new Promise((resolve, reject) => {
    server.once('error', reject)
    server.listen(port, '127.0.0.1', resolve)
  })
  const actual = server.address().port
  await new Promise(resolve => server.close(resolve))
  return actual
}
function httpCall(port, method, route, headers = {}, value) {
  return new Promise((resolve, reject) => {
    const req = http.request({hostname:'127.0.0.1', port, method, path:route, headers, agent:false, timeout:1500}, res => {
      let body = ''
      res.on('data', chunk => {body += chunk})
      res.on('error', reject)
      res.on('end', () => resolve({status:res.statusCode, body}))
    })
    req.on('error', reject)
    req.on('timeout', () => req.destroy(Error('HTTP smoke request timed out')))
    req.end(value === undefined ? undefined : JSON.stringify(value))
  })
}
async function waitForExit(child) {
  for (let i = 0; i < 50; i++) {
    if (child.exitCode !== null) return child.exitCode
    await delay(100)
  }
  throw Error('Child did not exit within five seconds')
}
async function main() {
  assert.equal(process.platform, 'win32', 'This smoke tool requires Windows')
  assert.ok(process.argv[2], 'Usage: node tools/smoke_listener.cjs <candidate.exe>')
  const exe = path.resolve(process.argv[2])
  assert.ok(fs.statSync(exe).isFile())
  // Refuse to run while any dashboard is listening. Verify fixture identity
  // again before mutations to protect against a process racing this probe.
  await freePort(17993)
  const port = await freePort(0)
  const cache = path.resolve(__dirname, '../.cache')
  fs.mkdirSync(cache, {recursive:true})
  const localData = fs.mkdtempSync(path.join(cache, 'listener-smoke-'))
  const settingsDir = path.join(localData, 'OpenCompanionListener')
  fs.mkdirSync(settingsDir)
  const settingsPath = path.join(settingsDir, 'settings.json')
  const key = crypto.randomBytes(32).toString('hex')
  const config = {listen_ip:'127.0.0.1', port, trusted_ip:'127.0.0.1', secret:key}
  fs.writeFileSync(settingsPath, JSON.stringify(config), {mode:0o600})
  const children = []
  const start = () => {
    const child = spawn(exe, [], {windowsHide:true, stdio:'ignore', env:{...process.env, LOCALAPPDATA:localData}})
    child.on('error', () => {child.smokeSpawnFailed = true})
    children.push(child)
    return child
  }
  const uiHeaders = {'X-OCL-Local-UI':'1', Origin:'http://127.0.0.1:17993'}
  const ui = (method, route, value) => httpCall(17993, method, route, uiHeaders, value)
  const child = start()
  try {
    let state
    for (let i = 0; i < 50; i++) {
      assert.ok(!child.smokeSpawnFailed, 'Could not launch candidate')
      try {
        const result = await ui('GET', '/ui/state')
        if (result.status === 200) {state = JSON.parse(result.body); break}
      } catch {}
      await delay(100)
    }
    assert.ok(state && state.config.secret === key, 'Fixture identity mismatch (key redacted)')
    assert.ok(state.running && !state.error, 'Candidate API did not start')
    assert.ok(!state.events.some(e => e.message.includes('Tray unavailable')), 'Native tray initialization failed')
    assert.equal((await httpCall(17993, 'GET', '/')).status, 200)
    assert.equal((await httpCall(17993, 'GET', '/ui/state')).status, 403)
    assert.equal((await httpCall(17993, 'GET', '/ui/state', {...uiHeaders, Host:'rebinding.example:17993'})).status, 403)
    const api = {host:'127.0.0.1', port}
    const status = await request(api, key, 'GET', '/v1/status')
    assert.equal(status.version, '0.1.3')
    const replayHeaders = makeHeaders(key, 'GET', '/v1/status', Buffer.alloc(0))
    assert.equal((await httpCall(port, 'GET', '/v1/status', replayHeaders)).status, 200)
    assert.equal((await ui('POST', '/ui/save', config)).status, 200)
    assert.equal((await httpCall(port, 'GET', '/v1/status', replayHeaders)).status, 401)
    assert.equal((await ui('POST', '/ui/restart')).status, 200)
    assert.equal((await httpCall(port, 'GET', '/v1/status', replayHeaders)).status, 401)
    assert.equal((await request(api, key, 'GET', '/v1/status')).ok, true)
    assert.equal(await waitForExit(start()), 0)
    assert.equal(child.exitCode, null, 'Original candidate unexpectedly exited')
    assert.equal((await ui('POST', '/ui/rotate')).status, 200)
    const disk = JSON.parse(fs.readFileSync(settingsPath, 'utf8'))
    assert.ok(disk.secret !== key, 'Key rotation did not persist (redacted)')
    await assert.rejects(request(api, key, 'GET', '/v1/status'), /HTTP 401/)
    assert.equal((await request(api, disk.secret, 'GET', '/v1/status')).ok, true)
    assert.equal((await ui('POST', '/ui/exit')).status, 200)
    assert.equal(await waitForExit(child), 0)
    console.log('PASS: candidate start, native tray initialization, UI guards, signed Node-to-Go status, save/restart replay, duplicate-instance exit, key rotation and clean shutdown')
    console.log('Not verified: visual tray/browser behavior, desktop input, sign-in startup, real Companion host or Windows 10/11 matrix')
  } finally {
    for (const process of children) if (process.exitCode === null && !process.killed) process.kill()
  }
}
main().catch(error => {console.error(error.message); process.exitCode = 1})
