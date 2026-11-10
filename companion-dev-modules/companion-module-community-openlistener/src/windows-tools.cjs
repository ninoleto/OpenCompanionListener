'use strict'
// URI IDs are official Windows ms-settings protocol values. Availability
// depends on Windows edition, build, hardware and policy; unsupported pages
// should produce a Windows shell error rather than launching a terminal.
const windowsSettingsChoices = Object.freeze([
  ['ms-settings:', 'General — Settings home'],
  ['ms-settings:display', 'System — Display'],
  ['ms-settings:display-advanced', 'System — Advanced display'],
  ['ms-settings:sound', 'System — Sound'],
  ['ms-settings:apps-volume', 'System — Volume mixer'],
  ['ms-settings:sound-devices', 'System — Audio devices'],
  ['ms-settings:notifications', 'System — Notifications'],
  ['ms-settings:powersleep', 'System — Power & sleep'],
  ['ms-settings:bluetooth', 'Devices — Bluetooth'],
  ['ms-settings:printers', 'Devices — Printers & scanners'],
  ['ms-settings:network-status', 'Network — Status'],
  ['ms-settings:network-wifi', 'Network — Wi-Fi'],
  ['ms-settings:network-ethernet', 'Network — Ethernet'],
  ['ms-settings:network-vpn', 'Network — VPN'],
  ['ms-settings:appsfeatures', 'Apps — Installed apps'],
  ['ms-settings:defaultapps', 'Apps — Default apps'],
  ['ms-settings:startupapps', 'Apps — Startup apps'],
  ['ms-settings:personalization', 'Personalization — Overview'],
  ['ms-settings:windowsupdate', 'Windows Update — Overview'],
  ['ms-settings:windowsdefender', 'Security — Windows Security']
].map(([id,label])=>({id,label})))
const controlPanelChoices = Object.freeze([
  ['appwiz.cpl','Programs and Features'],
  ['mmsys.cpl','Sound (classic)'],
  ['ncpa.cpl','Network Connections'],
  ['main.cpl','Mouse Properties'],
  ['desk.cpl','Display'],
  ['timedate.cpl','Date and Time'],
  ['intl.cpl','Region'],
  ['powercfg.cpl','Power Options'],
  ['sysdm.cpl','System Properties'],
  ['firewall.cpl','Windows Defender Firewall']
].map(([id,label])=>({id,label})))
module.exports = { windowsSettingsChoices, controlPanelChoices }
