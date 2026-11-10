'use strict'
const { test } = require('node:test')
const assert = require('node:assert/strict')
const { createHmac,createHash } = require('node:crypto')
const {makeHeaders} = require('./transport.cjs')

test('signatures match canonical HMAC payload',()=>{
 const key='ab'.repeat(32)
 const b=Buffer.from('{"type":"click","button":"left"}')
 const headers=makeHeaders(key,'POST','/v1/command',b,1700000000,'00'.repeat(16))
 const msg=['POST','/v1/command','1700000000','00'.repeat(16),createHash('sha256').update(b).digest('hex')].join('\n')
 assert.equal(headers['X-OCL-Signature'],createHmac('sha256',Buffer.from(key,'hex')).update(msg).digest('hex'))
})
test('rejects invalid keys',()=>assert.throws(()=>makeHeaders('bad','GET','/v1/status',Buffer.alloc(0))))
