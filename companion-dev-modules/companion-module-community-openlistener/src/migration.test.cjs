'use strict'
const test = require('node:test')
const assert = require('node:assert/strict')
const {upgradeScripts, shortcutFromText} = require('./upgrades.cjs')
const {presets} = require('./presets.cjs')
const {windowsSettingsChoices} = require('./windows-tools.cjs')
const {singleKeyChoices} = require('./keys.cjs')

function action(id,actionId,options){return {id,controlId:'bank:xyz',actionId,options}}
const migrate = (actions) => upgradeScripts[0]({currentConfig:null},{actions,feedbacks:[]})
test('migrate v0.1.2 builder, scroll, and text hotkey (raw values)',()=>{
  const acts=[
    action('id1','shortcut_builder',{ctrl:true,shift:false,alt:false,win:true,key:'F23'}),
    action('id2','scroll',{axis:'vertical',steps:-5}),
    action('id3','hotkey',{keys:'CTRL+SHIFT+F24'})
  ]
  const result=migrate(acts)
  assert.equal(result.updatedActions.length,3)
  assert.equal(acts[0].actionId,'hotkey')
  assert.equal(acts[0].options.mode,'builder')
  assert.equal(acts[0].options.win,true)
  assert.equal(acts[1].actionId,'scroll_slider')
  assert.deepEqual(acts[1].options,{axis:'vertical',magnitude:5,reverse:true})
  assert.equal(acts[2].actionId,'hotkey')
  assert.equal(acts[2].options.mode,'builder')
  assert.equal(acts[2].options.key,'F24')
  assert.equal(acts[2].options.ctrl,true)
  assert.equal(acts[2].options.shift,true)
  assert.deepEqual(acts.map(a=>a.id),['id1','id2','id3'])
})
test('retain complex key text, preserve wrapped types, and keep unrelated actions',()=>{
  const value=(v)=>({isExpression:false,value:v})
  const acts=[action('i','hotkey',{keys:value('CTRL+F23+F24')}),action('j','scroll',{axis:value('horizontal'),steps:value(10)}),action('k','open_file',{path:value('C:\\Windows\\notepad.exe')})]
  const result=migrate(acts)
  assert.equal(result.updatedActions.length,2)
  assert.equal(acts[0].options.mode.value,'advanced')
  assert.deepEqual(acts[0].options.keys,value('CTRL+F23+F24'))
  assert.equal(acts[1].options.magnitude.value,10)
  assert.equal(acts[1].options.reverse.value,false)
  assert.equal(acts[2].actionId,'open_file')
})
test('shortcuts parser rejects unsupported and ambiguous multiple base keys',()=>{
  assert.deepEqual(shortcutFromText('ALT+SHIFT+F13'),{ctrl:false,shift:true,alt:true,win:false,key:'F13'})
  assert.equal(shortcutFromText('CTRL+F23+F24'),null)
  assert.equal(shortcutFromText('CTRL+SUPER+Z'),null)
})
test('38 curated compatible presets, no unsupported actions or keys',()=>{
  assert.equal(Object.keys(presets).length,38)
  const supportedActions=new Set(['hotkey','scroll_slider','windows_settings'])
  const keys=new Set(singleKeyChoices.map(c=>c.id))
  const settings=new Set(windowsSettingsChoices.map(c=>c.id))
  for(const [id,p] of Object.entries(presets)){
    assert.equal(p.type,'button',id)
    assert.ok(p.category && p.name)
    assert.ok(p.style.color > 0 && p.style.bgcolor > 0)
    for(const a of p.steps[0].down){
      assert.ok(supportedActions.has(a.actionId),id)
      if(a.actionId==='hotkey')assert.ok(keys.has(a.options.key),id)
      if(a.actionId==='windows_settings')assert.ok(settings.has(a.options.page),id)
    }
  }
})
