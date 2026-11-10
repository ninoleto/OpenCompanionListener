'use strict'
const { test } = require('node:test')
const assert = require('node:assert/strict')
const { singleKeyChoices } = require('./keys.cjs')

test('single key dropdown contains every function key', () => {
  for (let n = 1; n <= 24; n++) assert.ok(singleKeyChoices.some(k => k.id === `F${n}`))
})
test('numpad keys are distinct from number-row and Enter keys', () => {
  for (let n = 0; n <= 9; n++) {
    assert.ok(singleKeyChoices.some(k => k.id === `NUMPAD${n}`))
    assert.ok(singleKeyChoices.some(k => k.id === String(n)))
  }
  assert.ok(singleKeyChoices.some(k => k.id === 'NUMPAD_ENTER'))
  assert.ok(singleKeyChoices.some(k => k.id === 'ENTER'))
})
test('arrow and navigation keys are present', () => {
  for (const id of ['UP','DOWN','LEFT','RIGHT','PGUP','PGDN','HOME','END','INSERT','DELETE']) {
    assert.ok(singleKeyChoices.some(k => k.id === id))
  }
})
test('dropdown IDs are unique and labels are meaningful', () => {
  const ids = singleKeyChoices.map(c => c.id)
  assert.equal(new Set(ids).size, ids.length)
  for (const choice of singleKeyChoices) {
    assert.ok(typeof choice.label === 'string' && choice.label.length > 4)
  }
})
