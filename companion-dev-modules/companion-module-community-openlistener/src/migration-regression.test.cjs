'use strict'
const test=require('node:test')
const assert=require('node:assert/strict')
const {runThroughUpgradeScripts}=require('@companion-module/base/dist/internal/upgrade.js')
const {upgradeScripts}=require('./upgrades.cjs')
const entry=require('@companion-module/base/dist/entrypoint.js')
const original=entry.runEntrypoint
let Class
entry.runEntrypoint=ctor=>{Class=ctor}
try {require('./main.cjs')} finally {entry.runEntrypoint=original}
function migrate(actionId,options,upgradeIndex=-1) {
  const action={id:'fixture',controlId:'bank:test',actionId,options:structuredClone(options),upgradeIndex}
  runThroughUpgradeScripts([action],[],upgradeIndex,upgradeScripts,{},{},false)
  return action
}
const literal=value=>({isExpression:false,value})
const expression=value=>({isExpression:true,value})
const unwrap=o=>Object.fromEntries(Object.entries(o).map(([k,v])=>[k,v&&typeof v==='object'?'value' in v?v.value:v:v]))
function runtime() {
  const calls=[]
  return {calls,actions:Class.prototype.actions.call({send:async c=>calls.push(c)})}
}
test('dynamic legacy scroll retains expression and has an executable migrated action',async()=>{
  for(const upgradeIndex of [-1,0]) {
    const axis=expression('$(custom:axis)'),steps=expression('$(custom:steps)')
    const migrated=migrate('scroll',{axis,steps},upgradeIndex)
    const {actions,calls}=runtime()
    assert.ok(actions[migrated.actionId])
    assert.deepEqual(migrated.options.steps,steps)
    assert.deepEqual(migrated.options.axis,axis)
    // Companion evaluates the preserved expressions before invoking callbacks.
    await actions[migrated.actionId].callback({options:{...unwrap(migrated.options),axis:'horizontal',steps:-7}})
    assert.deepEqual(calls,[{type:'scroll',axis:'horizontal',steps:-7}])
  }
})
test('literal-looking expression metadata survives shortcuts and scroll',()=>{
  for(const id of ['hotkey','key_down','key_up']) {
    for(const value of ['1','0','"CTRL+F23"','$(custom:keys)']) {
      const old=expression(value)
      assert.deepEqual(migrate(id,{keys:old}).options.keys,old)
    }
  }
  assert.deepEqual(migrate('scroll',{axis:literal('vertical'),steps:expression('-5')}).options.steps,expression('-5'))
})
test('unconvertible scroll fields remain editable and bounded at execution',async()=>{
  for(const value of ['',0,21,-21,1.5,'$(custom:steps)']) {
    const migrated=migrate('scroll',{axis:'vertical',steps:value})
    const {actions,calls}=runtime()
    assert.ok(actions[migrated.actionId])
    assert.equal(migrated.options.steps,value)
    await assert.rejects(actions[migrated.actionId].callback({options:unwrap(migrated.options)}))
    assert.equal(calls.length,0)
  }
})
test('builder expressions and unusual shortcut text survive migration',async()=>{
  const modifier=expression('$(custom:ctrl)')
  const built=migrate('shortcut_builder',{ctrl:modifier,shift:false,alt:false,win:true,key:expression('$(custom:key)')})
  assert.deepEqual(built.options.ctrl,modifier)
  assert.deepEqual(built.options.key,expression('$(custom:key)'))
  for(const keys of ['CTRL+F23+F24','RCTRL+NUMPAD_ENTER','CTRL+RETURN','', 'CTRL++A']) {
    const old=literal(keys)
    const migrated=migrate('hotkey',{keys:old})
    assert.deepEqual(migrated.options.keys,old)
  }
  const {actions,calls}=runtime()
  await actions.hotkey.callback({options:{mode:'advanced',keys:0}})
  assert.deepEqual(calls,[{type:'hotkey',keys:'0'}])
})
test('ordinary raw and wrapped keyboard/scroll migrations retain behavior and are idempotent',async()=>{
  for(const wrap of [v=>v,literal]) {
    for(const steps of [-20,-5,-1,1,5,20]) {
      const migrated=migrate('scroll',{axis:wrap('horizontal'),steps:wrap(steps)})
      const {actions,calls}=runtime()
      await actions[migrated.actionId].callback({options:unwrap(migrated.options)})
      assert.deepEqual(calls,[{type:'scroll',axis:'horizontal',steps}])
      const again=migrate(migrated.actionId,migrated.options,migrated.upgradeIndex)
      assert.deepEqual(again.options,migrated.options)
    }
    for(const id of ['hotkey','key_down','key_up']) {
      const migrated=migrate(id,{keys:wrap('CTRL+SHIFT+F23')})
      const {actions,calls}=runtime()
      await actions[migrated.actionId].callback({options:unwrap(migrated.options)})
      assert.deepEqual(calls,[{type:id,keys:'CTRL+SHIFT+F23'}])
    }
  }
})
test('existing v0.1.3 slider, file actions, and preset IDs are unchanged',()=>{
  for(const [id,options] of [['scroll_slider',{axis:'vertical',magnitude:5,reverse:true}],['open_file',{path:'C:\\Windows\\notepad.exe',args:''}]]) {
    assert.deepEqual(migrate(id,options,0).options,options)
  }
})
