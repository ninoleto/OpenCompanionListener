'use strict'

// Dropdown IDs are the native listener's stable key names. Never substitute a
// printable glyph for numpad keys: NUMPAD1 and 1 are different Windows keys.
const groups = [
  ['Function keys', Array.from({ length: 24 }, (_, i) => [`F${i + 1}`, `F${i + 1}`])],
  ['Navigation', [
    ['UP','Arrow Up'],['DOWN','Arrow Down'],['LEFT','Arrow Left'],['RIGHT','Arrow Right'],
    ['PGUP','Page Up'],['PGDN','Page Down'],['HOME','Home'],['END','End'],
    ['INSERT','Insert'],['DELETE','Delete']
  ]],
  ['Numpad', [
    ...Array.from({ length: 10 }, (_, i) => [`NUMPAD${i}`, `Numpad ${i}`]),
    ['NUMPAD_ENTER','Numpad Enter'],['NUMPAD_ADD','Numpad +'],
    ['NUMPAD_SUBTRACT','Numpad -'],['NUMPAD_MULTIPLY','Numpad *'],
    ['NUMPAD_DIVIDE','Numpad /'],['NUMPAD_DECIMAL','Numpad Decimal'],
    ['NUMPAD_SEPARATOR','Numpad Separator']
  ]],
  ['Common keys', [
    ['SPACE','Space'],['ENTER','Enter'],['ESC','Escape'],['TAB','Tab'],['BACKSPACE','Backspace'],
    ['CAPSLOCK','Caps Lock'],['SCROLLLOCK','Scroll Lock'],
    ['PRINTSCREEN','Print Screen'],['PAUSE','Pause / Break'],['APPS','Context Menu']
  ]],
  ['Modifiers', [
    ['SHIFT','Shift'],['CTRL','Ctrl'],['ALT','Alt'],['WIN','Windows'],
    ['LSHIFT','Left Shift'],['RSHIFT','Right Shift'],
    ['LCTRL','Left Ctrl'],['RCTRL','Right Ctrl'],
    ['LALT','Left Alt'],['RALT','Right Alt'],
    ['LWIN','Left Windows'],['RWIN','Right Windows']
  ]],
  ['Letters', Array.from('ABCDEFGHIJKLMNOPQRSTUVWXYZ', letter => [letter, letter])],
  ['Number row', Array.from('0123456789', digit => [digit, `Digit ${digit} (number row)`])],
  ['Punctuation (keyboard layout dependent)', [
    ['SEMICOLON','Semicolon'],['EQUALS','Equals / Plus key'],
    ['COMMA','Comma'],['MINUS','Minus'],['PERIOD','Period'],['SLASH','Slash'],
    ['BACKTICK','Backtick'],['LBRACKET','Left bracket'],['BACKSLASH','Backslash'],
    ['RBRACKET','Right bracket'],['APOSTROPHE','Apostrophe']
  ]],
  ['Media', [
    ['VOLUMEUP','Volume Up'],['VOLUMEDOWN','Volume Down'],['VOLUMEMUTE','Volume Mute'],
    ['MEDIA_PLAY_PAUSE','Play / Pause'],['MEDIA_NEXT','Next Track'],
    ['MEDIA_PREV','Previous Track'],['MEDIA_STOP','Stop Playback'],
    ['BROWSER_BACK','Browser Back'],['BROWSER_FORWARD','Browser Forward'],
    ['BROWSER_REFRESH','Browser Refresh'],['BROWSER_HOME','Browser Home']
  ]]
]

const singleKeyChoices = Object.freeze(groups.flatMap(([group, keys]) =>
  keys.map(([id, label]) => ({ id, label: `${label} — ${group}` }))
))

module.exports = { singleKeyChoices }
