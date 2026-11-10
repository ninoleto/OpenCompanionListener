'use strict'
// Companion Module API 1.x button presets. Users opt in by dragging these to a grid.
const {singleKeyChoices} = require('./keys.cjs')
const available = new Set(singleKeyChoices.map(k=>k.id))
const colors = {Windows:0x233e59, Chrome:0x284a51, Editing:0x403358, Mouse:0x254d46, Settings:0x3b3e59}
const presets={}
function add(id,category,name,shortcut,options={}) {
  const actionId = options.actionId || 'hotkey'
  const actionOptions = options.payload || shortcutOptions(shortcut)
  const words = name.replace(/ /g,'\n')
  presets[id]={
    type:'button',category,name,
    style:{text: words.length > 28 ? name : words, size:'auto',color:0xffffff,bgcolor:colors[category]},
    steps:[{down:[{actionId,options:actionOptions}],up:[]}],feedbacks:[]
  }
}
function shortcutOptions(shortcut){
  const tokens=shortcut.split('+').map(x=>x.trim().toUpperCase())
  const key=tokens.at(-1)
  if(!available.has(key))throw new Error('Preset uses unsupported key '+key)
  return {mode:'builder',key,ctrl:tokens.includes('CTRL'),shift:tokens.includes('SHIFT'),alt:tokens.includes('ALT'),win:tokens.includes('WIN')}
}
// Windows
for(const [id,name,combo] of [
 ['win_explorer','File Explorer','WIN+E'],['win_desktop','Show Desktop','WIN+D'],
 ['win_settings','Windows Settings','WIN+I'],['win_task_manager','Task Manager','CTRL+SHIFT+ESC'],
 ['win_task_view','Task View','WIN+TAB'],['win_capture','Screen Snip','WIN+SHIFT+S'],
 ['win_run','Run Dialog','WIN+R'],['win_clipboard','Clipboard History','WIN+V'],
 ['win_snap_left','Snap Left','WIN+LEFT'],['win_snap_right','Snap Right','WIN+RIGHT'],
 ['win_project','Project Display','WIN+P'],['win_emojis','Emoji Picker','WIN+PERIOD']
])add(id,'Windows',name,combo)
// Chrome, applied to the currently focused Chrome window (or browser with matching shortcuts)
for(const [id,name,combo] of [
 ['chrome_new_tab','New Tab','CTRL+T'],['chrome_close_tab','Close Tab','CTRL+W'],
 ['chrome_restore_tab','Reopen Tab','CTRL+SHIFT+T'],['chrome_next_tab','Next Tab','CTRL+TAB'],
 ['chrome_prev_tab','Previous Tab','CTRL+SHIFT+TAB'],['chrome_address','Address Bar','CTRL+L'],
 ['chrome_downloads','Downloads','CTRL+J'],['chrome_history','History','CTRL+H'],
 ['chrome_incognito','Incognito','CTRL+SHIFT+N'],['chrome_refresh','Refresh','F5'],
 ['chrome_find','Find in Page','CTRL+F'],['chrome_new_window','New Window','CTRL+N']
])add(id,'Chrome',name,combo)
// Editing
for(const [id,name,combo] of [
 ['edit_copy','Copy','CTRL+C'],['edit_paste','Paste','CTRL+V'],['edit_cut','Cut','CTRL+X'],
 ['edit_undo','Undo','CTRL+Z'],['edit_redo','Redo','CTRL+Y'],
 ['edit_select_all','Select All','CTRL+A'],['edit_save','Save','CTRL+S'],
 ['edit_find','Find','CTRL+F']
])add(id,'Editing',name,combo)
// Mouse wheel presets
for(const [id,name,steps] of [['wheel_up','Scroll Up',1],['wheel_down','Scroll Down',-1]])
  add(id,'Mouse',name,'',{actionId:'scroll_slider',payload:{axis:'vertical',magnitude:Math.abs(steps),reverse:steps<0}})
for(const [id,name,page] of [
 ['settings_sound','Sound Settings','ms-settings:sound'],
 ['settings_volume','Volume Mixer','ms-settings:apps-volume'],
 ['settings_display','Display Settings','ms-settings:display'],
 ['settings_network','Network Settings','ms-settings:network-status']
])add(id,'Settings',name,'',{actionId:'windows_settings',payload:{page}})

module.exports={presets}
