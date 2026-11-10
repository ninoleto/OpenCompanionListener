package main

import (
	"strconv"
	"strings"
)

// keyDef describes the Windows virtual key and whether its scan code must
// carry KEYEVENTF_EXTENDEDKEY. This distinction matters for the navigation
// cluster, right-side Ctrl/Alt, NumPad Enter, and NumPad Divide.
type keyDef struct {
	VK       uint16
	Extended bool
	Scan     uint16
}

func keySpec(s string) (keyDef, bool) {
	s = strings.ToUpper(strings.TrimSpace(s))
	keys := map[string]keyDef{
		"CTRL": {VK: 0x11}, "CONTROL": {VK: 0x11}, "ALT": {VK: 0x12}, "SHIFT": {VK: 0x10},
		"WIN": {VK: 0x5B, Extended: true}, "WINDOWS": {VK: 0x5B, Extended: true},
		"LWIN": {VK: 0x5B, Extended: true}, "RWIN": {VK: 0x5C, Extended: true},
		"LCTRL": {VK: 0xA2}, "RCTRL": {VK: 0xA3, Extended: true},
		"LSHIFT": {VK: 0xA0}, "RSHIFT": {VK: 0xA1},
		"LALT": {VK: 0xA4}, "RALT": {VK: 0xA5, Extended: true},
		"ENTER": {VK: 0x0D}, "RETURN": {VK: 0x0D},
		"ESC": {VK: 0x1B}, "ESCAPE": {VK: 0x1B},
		"TAB": {VK: 0x09}, "SPACE": {VK: 0x20}, "BACKSPACE": {VK: 0x08},
		"DELETE": {VK: 0x2E, Extended: true}, "DEL": {VK: 0x2E, Extended: true},
		"INSERT": {VK: 0x2D, Extended: true}, "INS": {VK: 0x2D, Extended: true},
		"HOME": {VK: 0x24, Extended: true}, "END": {VK: 0x23, Extended: true},
		"PGUP": {VK: 0x21, Extended: true}, "PAGEUP": {VK: 0x21, Extended: true}, "PAGE_UP": {VK: 0x21, Extended: true},
		"PGDN": {VK: 0x22, Extended: true}, "PAGEDOWN": {VK: 0x22, Extended: true}, "PAGE_DOWN": {VK: 0x22, Extended: true},
		"UP": {VK: 0x26, Extended: true}, "DOWN": {VK: 0x28, Extended: true},
		"LEFT": {VK: 0x25, Extended: true}, "RIGHT": {VK: 0x27, Extended: true},
		"CAPSLOCK": {VK: 0x14}, "PRINTSCREEN": {VK: 0x2C, Extended: true},
		"NUMLOCK": {VK: 0x90, Extended: true, Scan: 0x45}, "SCROLLLOCK": {VK: 0x91},
		"PAUSE": {VK: 0x13}, "APPS": {VK: 0x5D, Extended: true}, "MENU": {VK: 0x5D, Extended: true},
		"NUMPAD_ADD": {VK: 0x6B}, "NUMPAD_PLUS": {VK: 0x6B},
		"NUMPAD_SUBTRACT": {VK: 0x6D}, "NUMPAD_MINUS": {VK: 0x6D},
		"NUMPAD_MULTIPLY": {VK: 0x6A}, "NUMPAD_DIVIDE": {VK: 0x6F, Extended: true},
		"NUMPAD_DECIMAL": {VK: 0x6E}, "NUMPAD_SEPARATOR": {VK: 0x6C},
		"NUMPAD_ENTER": {VK: 0x0D, Extended: true}, "NUMPADENTER": {VK: 0x0D, Extended: true},
		"VOLUMEUP": {VK: 0xAF, Extended: true}, "VOLUMEDOWN": {VK: 0xAE, Extended: true},
		"VOLUMEMUTE":       {VK: 0xAD, Extended: true},
		"MEDIA_PLAY_PAUSE": {VK: 0xB3, Extended: true}, "MEDIA_NEXT": {VK: 0xB0, Extended: true},
		"MEDIA_PREV": {VK: 0xB1, Extended: true}, "MEDIA_STOP": {VK: 0xB2, Extended: true},
		"BROWSER_BACK": {VK: 0xA6, Extended: true}, "BROWSER_FORWARD": {VK: 0xA7, Extended: true},
		"BROWSER_REFRESH": {VK: 0xA8, Extended: true}, "BROWSER_HOME": {VK: 0xAC, Extended: true},
		"SEMICOLON": {VK: 0xBA}, "EQUALS": {VK: 0xBB}, "COMMA": {VK: 0xBC},
		"MINUS": {VK: 0xBD}, "PERIOD": {VK: 0xBE}, "SLASH": {VK: 0xBF},
		"BACKTICK": {VK: 0xC0}, "LBRACKET": {VK: 0xDB}, "BACKSLASH": {VK: 0xDC},
		"RBRACKET": {VK: 0xDD}, "APOSTROPHE": {VK: 0xDE},
	}
	if k, ok := keys[s]; ok {
		return k, true
	}
	if len(s) == 1 && ((s[0] >= 'A' && s[0] <= 'Z') || (s[0] >= '0' && s[0] <= '9')) {
		return keyDef{VK: uint16(s[0])}, true
	}
	if strings.HasPrefix(s, "F") {
		n, err := strconv.Atoi(s[1:])
		if err == nil && n >= 1 && n <= 24 {
			return keyDef{VK: uint16(0x70 + n - 1)}, true
		}
	}
	if strings.HasPrefix(s, "NUMPAD") {
		n, err := strconv.Atoi(strings.TrimPrefix(s, "NUMPAD"))
		if err == nil && n >= 0 && n <= 9 {
			return keyDef{VK: uint16(0x60 + n)}, true
		}
	}
	return keyDef{}, false
}

// Retained for compatibility with existing validation/tests.
func keyCode(s string) (uint16, bool) {
	k, ok := keySpec(s)
	return k.VK, ok
}
