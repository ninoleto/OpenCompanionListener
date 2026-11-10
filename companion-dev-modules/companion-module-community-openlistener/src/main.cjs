'use strict'
const { InstanceBase, InstanceStatus, runEntrypoint } = require('@companion-module/base')
const { request } = require('./transport.cjs')
const { singleKeyChoices } = require('./keys.cjs')
const { windowsSettingsChoices, controlPanelChoices } = require('./windows-tools.cjs')
const { presets } = require('./presets.cjs')
const { upgradeScripts } = require('./upgrades.cjs')

class OpenListener extends InstanceBase {
  constructor(internal) { super(internal); this.config = {}; this.key = ''; this.monitor = null; this.online = false }

  async init(config, _firstInit, secrets) {
    this.setActionDefinitions(this.actions())
    this.setPresetDefinitions(presets)
    await this.configUpdated(config, secrets)
  }
  async configUpdated(config, secrets) {
    this.config = config || {}
    this.key = secrets?.pairing_key || ''
    this.online = false
    if (this.monitor) clearInterval(this.monitor)
    this.updateStatus(InstanceStatus.Connecting)
    void this.ping()
    this.monitor = setInterval(() => { void this.ping() }, 6500)
  }
  async destroy() { if (this.monitor) clearInterval(this.monitor); this.monitor = null }
  getConfigFields() {
    return [
      { type: 'textinput', id: 'host', label: 'Listener IP address', width: 8, default: '127.0.0.1' },
      { type: 'number', id: 'port', label: 'Listener port', width: 4, default: 17992, min: 1024, max: 65535 },
      { type: 'secret-text', id: 'pairing_key', label: 'Pairing key (from listener Settings)', width: 12 }
    ]
  }
  async ping() {
    try {
      await request(this.config, this.key, 'GET', '/v1/status')
      if (!this.online) { this.online = true; this.updateStatus(InstanceStatus.Ok) }
    } catch (err) {
      if (this.online) this.log('warn', `Listener connection lost: ${err.message}`)
      this.online = false
      this.updateStatus(InstanceStatus.ConnectionFailure, 'Cannot reach or authenticate with listener')
    }
  }
  async send(cmd) {
    try {
      await request(this.config, this.key, 'POST', '/v1/command', cmd)
      if (!this.online) {this.online = true; this.updateStatus(InstanceStatus.Ok)}
    } catch(err) {
      this.log('error', `Action ${cmd.type} failed: ${err.message}`)
      throw err
    }
  }
  actions() {
    const text = (id, label, value) => ({ type:'textinput', id, label, default: value })
    const number = (id, label, value) => ({ type:'number', id, label, default:value, min:-10000, max:10000 })
    const button = {type:'dropdown', id:'button', label:'Mouse button', default:'left', choices:[{id:'left',label:'Left'},{id:'right',label:'Right'},{id:'middle',label:'Middle'}]}
    const actions = {}
    // Dedicated single-key actions for encoders and regular buttons. The
    // original free-form shortcut actions remain available and unchanged.
    const singleKeyOption = {
      type:'dropdown', id:'key', label:'Key', default:'ENTER', choices:singleKeyChoices
    }
    for (const [id,name,type] of [
      ['single_key','Press single key','hotkey'],
      ['single_key_down','Hold single key','key_down'],
      ['single_key_up','Release single key','key_up']
    ]) {
      actions[id] = {
        name, options:[singleKeyOption],
        callback:async(ev)=>this.send({type,keys:String(ev.options.key)})
      }
    }
    for (const [id,name] of [['click','Mouse click'],['mouse_down','Mouse button down'],['mouse_up','Mouse button up']]) {
      actions[id] = { name, options:[button], callback: async (ev)=>this.send({type:id,button:ev.options.button}) }
    }
    // Builder is the primary interface for press, hold, and release.
    // Advanced mode retains uncommon legacy multi-key expressions.
    const comboOptions = [
      {type:'dropdown',id:'mode',label:'Shortcut configuration',default:'builder',choices:[
        {id:'builder',label:'Checkbox + key dropdown'}, {id:'advanced',label:'Advanced / existing shortcut text'}]},
      ...[['ctrl','Ctrl'],['shift','Shift'],['alt','Alt'],['win','Windows']].map(([id,label]) => ({
        type:'checkbox',id,label,default:false,isVisibleExpression:'$(options:mode) == "builder"'
      })),
      {type:'dropdown',id:'key',label:'Key',default:'F23',choices:singleKeyChoices,isVisibleExpression:'$(options:mode) == "builder"'},
      {type:'textinput',id:'keys',label:'Advanced shortcut',default:'',isVisibleExpression:'$(options:mode) == "advanced"'}
    ]
    for(const [id,name] of [['hotkey','Press keyboard shortcut'],['key_down','Hold keyboard shortcut'],['key_up','Release keyboard shortcut']]){
      actions[id]={name,options:comboOptions,callback: async ev=>{
        const o=ev.options
        const raw=String(o.keys||'').trim()
        // Existing saved configurations may predate the new mode field.
        const useAdvanced=o.mode==='advanced' || (!o.mode && raw)
        const keys=useAdvanced ? raw : [o.ctrl&&'CTRL',o.shift&&'SHIFT',o.alt&&'ALT',o.win&&'WIN',String(o.key||'F23')].filter(Boolean).join('+')
        if(!keys)throw new Error('Shortcut is empty')
        return this.send({type:id,keys})
      }}
    }
    actions.num_lock = {
      name:'Num Lock — Toggle / ON / OFF',
      options:[{type:'dropdown',id:'mode',label:'Num Lock state',default:'toggle',choices:[
        {id:'toggle',label:'Toggle state'}, {id:'on',label:'Turn ON'}, {id:'off',label:'Turn OFF'}
      ]}],
      callback: async(ev)=>this.send({type:'num_lock',mode:String(ev.options.mode)})
    }
    actions.scroll_slider = {
      name:'Mouse wheel scroll',
      options:[
        {type:'dropdown',id:'axis',label:'Axis',default:'vertical',choices:[{id:'vertical',label:'Vertical'},{id:'horizontal',label:'Horizontal'}]},
        {type:'number',id:'magnitude',label:'Scroll steps',default:1,min:1,max:20,step:1,range:true},
        {type:'checkbox',id:'reverse',label:'Reverse direction (negative)',default:false}
      ],
      callback:async(ev)=>{
        const raw = Number(ev.options.magnitude)
        if (!Number.isSafeInteger(raw) || raw < 1 || raw > 20) throw Error('Scroll steps must be between 1 and 20')
        return this.send({type:'scroll',axis:ev.options.axis,steps:ev.options.reverse ? -raw : raw})
      }
    }
    actions.windows_settings = {
      name:'Windows 10/11 Settings (dropdown)',
      options:[{type:'dropdown',id:'page',label:'Settings page',default:'ms-settings:',choices:windowsSettingsChoices}],
      callback:async(ev)=>this.send({type:'open_file',path:String(ev.options.page),args:''})
    }
    actions.control_panel = {
      name:'Windows Control Panel (dropdown)',
      options:[{type:'dropdown',id:'tool',label:'Control Panel tool',default:'appwiz.cpl',choices:controlPanelChoices}],
      callback:async(ev)=>this.send({type:'open_file',path:'control.exe',args:String(ev.options.tool)})
    }
    actions.move = {name:'Move mouse (relative)', options:[number('dx','Horizontal pixels',0),number('dy','Vertical pixels',0)],callback:async(ev)=>this.send({type:'move',dx:Number(ev.options.dx),dy:Number(ev.options.dy)})}
    actions.open_url = {name:'Open URL in default browser', options:[text('url','Web address','https://example.com/')], callback:async(ev)=>this.send({type:'open_url',url:String(ev.options.url)})}
    actions.open_file = {name:'Open application, file or folder', options:[text('path','Absolute path','C:\\Windows\\notepad.exe'),text('args','Arguments (optional)','')], callback:async(ev)=>this.send({type:'open_file',path:String(ev.options.path),args:String(ev.options.args||'')})}
    return actions
  }
}
runEntrypoint(OpenListener, upgradeScripts)
