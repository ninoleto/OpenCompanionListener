'use strict'
// API 1.x module upgrade script. Run once by Companion for existing actions.
// Never delete scripts that have been published: Companion indexes them.
const MODIFIERS = Object.freeze({CTRL:'ctrl',CONTROL:'ctrl',SHIFT:'shift',ALT:'alt',WIN:'win',WINDOWS:'win'})
function optionValue(value) {
  return value && typeof value === 'object' && Object.prototype.hasOwnProperty.call(value,'value') ? value.value : value
}
function sameWrap(value, template) {
  return template && typeof template === 'object' && Object.prototype.hasOwnProperty.call(template,'value')
    ? {isExpression:false,value} : value
}
function shortcutFromText(text) {
  const parts = String(text || '').toUpperCase().split('+').map(s=>s.trim()).filter(Boolean)
  if(!parts.length) return null
  const state = {ctrl:false,shift:false,alt:false,win:false}
  for (const part of parts.slice(0,-1)) {
    const flag=MODIFIERS[part]
    if (!flag || state[flag]) return null
    state[flag]=true
  }
  // Match the supported dropdown keys by ID. Unknown/custom combinations use advanced mode.
  const {singleKeyChoices} = require('./keys.cjs')
  const key=parts.at(-1)
  if(!singleKeyChoices.some(c=>c.id===key)) return null
  return {...state,key}
}
function migrateExisting(_context, props) {
  const updatedActions=[]
  for (const action of props.actions || []) {
    if(!action || !action.options)continue
    const options=action.options
    const sample=Object.values(options)[0]
    const put=(key,val)=>sameWrap(val,options[key]??sample)
    let change=false
    if(action.actionId==='shortcut_builder') {
      action.actionId='hotkey'
      action.options={mode:put('mode','builder'),key:options.key,...Object.fromEntries(['ctrl','shift','alt','win'].map(k=>[k,options[k]??put(k,false)]))}
      change=true
    } else if(['hotkey','key_down','key_up'].includes(action.actionId)) {
      if (Object.prototype.hasOwnProperty.call(options,'keys') && !Object.prototype.hasOwnProperty.call(options,'mode')) {
        const raw=optionValue(options.keys)
        const parsed=typeof raw==='string'? shortcutFromText(raw):null
        if(parsed){
          action.options={mode:put('mode','builder'),...Object.fromEntries(Object.entries(parsed).map(([k,v])=>[k,put(k,v)]))}
        } else {
          // Complex and expression shortcuts MUST be preserved verbatim.
          action.options={...options,mode:put('mode','advanced')}
        }
        change=true
      }
    } else if(action.actionId==='scroll') {
      const value=Number(optionValue(options.steps))
      if(Number.isSafeInteger(value) && value!==0 && Math.abs(value)<=20){
        action.actionId='scroll_slider'
        action.options={axis:options.axis??put('axis','vertical'),magnitude:put('magnitude',Math.abs(value)),reverse:put('reverse',value<0)}
        change=true
      }
      // If an old value cannot convert safely, leave it untouched.
    }
    if(change)updatedActions.push(action)
  }
  return {updatedConfig:null,updatedSecrets:null,updatedActions,updatedFeedbacks:[]}
}
module.exports={upgradeScripts:[migrateExisting],shortcutFromText}
