'use strict'
const assert = require('node:assert/strict')
const test = require('node:test')
const Module = require('node:module')
const { singleKeyChoices } = require('./keys.cjs')
const { windowsSettingsChoices, controlPanelChoices } = require('./windows-tools.cjs')

// Capture module class without connecting to Companion; no network is used.
const originalLoad = Module._load
let Class
let RegisteredUpgradeScripts
Module._load = function(id, parent, isMain) {
  if (id === '@companion-module/base') {
    return {
      InstanceBase: class {},
      InstanceStatus: {Connecting: 1, Ok: 2, ConnectionFailure: 3},
      runEntrypoint: (ctor, scripts) => {Class = ctor; RegisteredUpgradeScripts = scripts}
    }
  }
  return originalLoad.call(this, id, parent, isMain)
}
try {require('./main.cjs')} finally {Module._load = originalLoad}
function fixture() {
  const app = new Class()
  const calls = []
  app.send = async cmd => {calls.push(cmd);return {ok:true}}
  return {actions:app.actions(),calls}
}

test('Num Lock only in dedicated dropdown, never in single key choices', () => {
  assert.ok(!singleKeyChoices.some(x=>x.id==='NUMLOCK'))
  const {actions,calls}=fixture()
  assert.ok(actions.single_key)
  assert.ok(actions.single_key_down)
  assert.ok(actions.single_key_up)
  assert.ok(actions.num_lock)
  return actions.num_lock.callback({options:{mode:'off'}}).then(()=>{
    assert.deepEqual(calls.at(-1),{type:'num_lock',mode:'off'})
  })
})

test('only builder UI is exposed; legacy actions are migrated, not duplicated', async () => {
  const {actions,calls}=fixture()
  for(const id of ['hotkey','key_down','key_up','scroll_slider','open_url','open_file','click','mouse_down','mouse_up','move']) assert.ok(actions[id],id)
  for(const id of ['shortcut_builder','scroll']) assert.equal(actions[id],undefined,id)
  await actions.hotkey.callback({options:{mode:'builder',ctrl:true,shift:true,alt:false,win:false,key:'F23'}})
  await actions.key_down.callback({options:{mode:'builder',ctrl:true,shift:false,alt:false,win:false,key:'F24'}})
  await actions.key_up.callback({options:{mode:'builder',ctrl:true,shift:false,alt:false,win:false,key:'F24'}})
  assert.deepEqual(calls,[
    {type:'hotkey',keys:'CTRL+SHIFT+F23'},
    {type:'key_down',keys:'CTRL+F24'},
    {type:'key_up',keys:'CTRL+F24'}
  ])
})

test('legacy unconvertible multi-key combo remains executable in advanced mode', async()=>{
  const {actions,calls}=fixture()
  await actions.hotkey.callback({options:{mode:'advanced',keys:'CTRL+F23+F24'}})
  await actions.key_up.callback({options:{keys:'CTRL+F23+F24'}})
  assert.deepEqual(calls,[{type:'hotkey',keys:'CTRL+F23+F24'}, {type:'key_up',keys:'CTRL+F23+F24'}])
})

test('scroll slider emits one signed integer and rejects out-of-range input', async()=>{
  const {actions,calls}=fixture()
  const action=actions.scroll_slider
  const step=action.options.find(x=>x.id==='magnitude')
  assert.equal(step.range,true)
  assert.equal(step.min,1)
  assert.equal(step.max,20)
  await action.callback({options:{axis:'horizontal',magnitude:5,reverse:true}})
  assert.deepEqual(calls.at(-1),{type:'scroll',axis:'horizontal',steps:-5})
  await assert.rejects(()=>action.callback({options:{axis:'vertical',magnitude:100,reverse:false}}))
})

test('Windows Settings and Control Panel selectors contain expected tools',async()=>{
  const {actions,calls}=fixture()
  assert.ok(windowsSettingsChoices.some(x=>x.id==='ms-settings:apps-volume'))
  assert.ok(controlPanelChoices.some(x=>x.id==='appwiz.cpl'))
  await actions.windows_settings.callback({options:{page:'ms-settings:apps-volume'}})
  await actions.control_panel.callback({options:{tool:'appwiz.cpl'}})
  assert.deepEqual(calls,[{type:'open_file',path:'ms-settings:apps-volume',args:''},{type:'open_file',path:'control.exe',args:'appwiz.cpl'}])
})

test('module registers presets after action definitions and registers upgrade migration', async()=>{
  assert.equal(RegisteredUpgradeScripts.length,1)
  const inst=new Class()
  const ordered=[]
  inst.setActionDefinitions=defs=>{ordered.push('actions');assert.ok(defs.hotkey)}
  inst.setPresetDefinitions=defs=>{ordered.push('presets');assert.equal(Object.keys(defs).length,38)}
  inst.configUpdated=async()=>{ordered.push('config')}
  await inst.init({},false,{})
  assert.deepEqual(ordered,['actions','presets','config'])
})
