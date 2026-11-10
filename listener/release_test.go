package main

import "testing"

func TestNumLockCommandAndScan(t *testing.T) {
	key, ok := keySpec("NUMLOCK")
	if !ok || key.VK != 0x90 || key.Scan != 0x45 || !key.Extended {
		t.Fatalf("NumLock input flags not configured: %+v / %v", key, ok)
	}
	for _, mode := range []string{"toggle", "on", "off"} {
		if err := validateCommand(Command{Type: "num_lock", Mode: mode}); err != nil {
			t.Fatalf("valid mode %s rejected: %v", mode, err)
		}
	}
	for _, mode := range []string{"", "true", "disable", "toggle; exit"} {
		if err := validateCommand(Command{Type: "num_lock", Mode: mode}); err == nil {
			t.Errorf("invalid mode %q accepted", mode)
		}
	}
}

func TestScrollLimitsAndLegacyActions(t *testing.T) {
	for _, n := range []int{-20, -10, -1, 1, 5, 10, 20} {
		if err := validateCommand(Command{Type: "scroll", Axis: "vertical", Steps: n}); err != nil {
			t.Errorf("scroll %d rejected: %v", n, err)
		}
	}
	for _, n := range []int{0, 21, -21, 999999} {
		if err := validateCommand(Command{Type: "scroll", Axis: "vertical", Steps: n}); err == nil {
			t.Errorf("scroll %d accepted", n)
		}
	}
	for _, cmd := range []Command{
		{Type: "hotkey", Keys: "CTRL+SHIFT+F23"},
		{Type: "key_down", Keys: "NUMPAD_ENTER"},
		{Type: "key_up", Keys: "NUMPAD_ENTER"},
		{Type: "open_file", Path: "ms-settings:sound"},
		{Type: "open_file", Path: "control.exe", Args: "appwiz.cpl"},
		{Type: "open_file", Path: `C:\Windows\notepad.exe`},
		{Type: "open_url", URL: "https://web.whatsapp.com/"},
	} {
		if err := validateCommand(cmd); err != nil {
			t.Errorf("legacy/new command %+v rejected: %v", cmd, err)
		}
	}
}
