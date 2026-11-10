package main

import "testing"

func TestSingleKeySupport(t *testing.T) {
	checks := []struct {
		name string
		vk   uint16
		ext  bool
	}{
		{"F13", 0x7C, false}, {"F24", 0x87, false},
		{"NUMPAD0", 0x60, false}, {"NUMPAD9", 0x69, false},
		{"NUMPAD_ADD", 0x6B, false}, {"NUMPAD_DIVIDE", 0x6F, true},
		{"NUMPAD_ENTER", 0x0D, true}, {"ENTER", 0x0D, false},
		{"LEFT", 0x25, true}, {"UP", 0x26, true},
		{"PGUP", 0x21, true}, {"PGDN", 0x22, true},
		{"HOME", 0x24, true}, {"END", 0x23, true},
		{"DELETE", 0x2E, true}, {"INSERT", 0x2D, true},
		{"RCTRL", 0xA3, true}, {"LCTRL", 0xA2, false},
		{"SPACE", 0x20, false}, {"A", 0x41, false},
	}
	for _, tc := range checks {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := keySpec(tc.name)
			if !ok || got.VK != tc.vk || got.Extended != tc.ext {
				t.Fatalf("%s: got %+v, %v; expected vk=%02X extended=%v", tc.name, got, ok, tc.vk, tc.ext)
			}
			if err := validateCommand(Command{Type: "hotkey", Keys: tc.name}); err != nil {
				t.Fatal(err)
			}
		})
	}
	for _, name := range []string{"F0", "F25", "NUMPAD10", "NUMPAD_7", "NOT_A_KEY"} {
		if _, ok := keySpec(name); ok {
			t.Errorf("unexpectedly accepted %q", name)
		}
	}
}
