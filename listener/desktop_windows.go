//go:build windows

package main

import (
	"errors"
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

var u32 = syscall.NewLazyDLL("user32.dll")
var shell32 = syscall.NewLazyDLL("shell32.dll")
var sendInput = u32.NewProc("SendInput")
var shellExecute = shell32.NewProc("ShellExecuteW")
var getKeyState = u32.NewProc("GetKeyState")

// INPUT for 64-bit Windows: 8-byte type/padding followed by 32-byte union.
type winInput struct {
	Type  uint32
	Pad   uint32
	A     uint32
	B     uint32
	C     uint32
	D     uint32
	E     uint32
	Pad2  uint32
	Extra uintptr
}

const (
	mouseMove       = 0x0001
	mouseLeftDown   = 0x0002
	mouseLeftUp     = 0x0004
	mouseRightDown  = 0x0008
	mouseRightUp    = 0x0010
	mouseMiddleDown = 0x0020
	mouseMiddleUp   = 0x0040
	mouseWheel      = 0x0800
	mouseHWheel     = 0x1000
	keyExtended     = 0x0001
	keyUp           = 0x0002
)

func putInput(in winInput) error {
	if unsafe.Sizeof(in) != 40 {
		return errors.New("unsupported architecture (expected amd64)")
	}
	r, _, err := sendInput.Call(1, uintptr(unsafe.Pointer(&in)), uintptr(unsafe.Sizeof(in)))
	if r != 1 {
		return fmt.Errorf("SendInput failed: %v", err)
	}
	return nil
}
func sendKey(key keyDef, up bool) error {
	var flags uint32
	if key.Extended {
		flags |= keyExtended
	}
	if up {
		flags |= keyUp
	}
	in := winInput{Type: 1, A: uint32(key.VK) | uint32(key.Scan)<<16, B: flags} // A low 16: wVk; B at offset 12: dwFlags
	return putInput(in)
}
func sendMouse(flags uint32, dx, dy int32, data uint32) error {
	return putInput(winInput{A: uint32(dx), B: uint32(dy), C: data, D: flags})
}
func mouseFlags(button string) (uint32, uint32) {
	switch button {
	case "right":
		return mouseRightDown, mouseRightUp
	case "middle":
		return mouseMiddleDown, mouseMiddleUp
	default:
		return mouseLeftDown, mouseLeftUp
	}
}
func shellOpen(target, args string) error {
	verb, _ := syscall.UTF16PtrFromString("open")
	path, err := syscall.UTF16PtrFromString(target)
	if err != nil {
		return err
	}
	var params *uint16
	if args != "" {
		params, err = syscall.UTF16PtrFromString(args)
		if err != nil {
			return err
		}
	}
	r, _, _ := shellExecute.Call(0, uintptr(unsafe.Pointer(verb)), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(params)), 0, 1)
	if r <= 32 {
		return fmt.Errorf("Windows ShellExecuteW returned %d", r)
	}
	return nil
}
func executeDesktop(c Command) error {
	switch c.Type {
	case "num_lock":
		state, _, _ := getKeyState.Call(uintptr(0x90))
		enabled := state&1 != 0
		if (c.Mode == "on" && enabled) || (c.Mode == "off" && !enabled) {
			return nil
		}
		k, _ := keySpec("NUMLOCK")
		if err := sendKey(k, false); err != nil {
			return err
		}
		return sendKey(k, true)
	case "open_url":
		return shellOpen(c.URL, "")
	case "open_file":
		return shellOpen(c.Path, c.Args)
	case "hotkey", "key_down", "key_up":
		ks := strings.Split(c.Keys, "+")
		var codes []keyDef
		for _, k := range ks {
			key, _ := keySpec(k)
			codes = append(codes, key)
		}
		if c.Type == "key_up" {
			for i := len(codes) - 1; i >= 0; i-- {
				if err := sendKey(codes[i], true); err != nil {
					return err
				}
			}
			return nil
		}
		for i, vk := range codes {
			if err := sendKey(vk, false); err != nil {
				for j := i - 1; j >= 0; j-- {
					_ = sendKey(codes[j], true)
				}
				return err
			}
		}
		if c.Type == "key_down" {
			return nil
		}
		for i := len(codes) - 1; i >= 0; i-- {
			if err := sendKey(codes[i], true); err != nil {
				return err
			}
		}
		return nil
	case "click", "mouse_down", "mouse_up":
		down, up := mouseFlags(c.Button)
		if c.Type != "mouse_up" {
			if err := sendMouse(down, 0, 0, 0); err != nil {
				return err
			}
		}
		if c.Type != "mouse_down" {
			if err := sendMouse(up, 0, 0, 0); err != nil {
				return err
			}
		}
		return nil
	case "move":
		return sendMouse(mouseMove, int32(c.DX), int32(c.DY), 0)
	case "scroll":
		flag := uint32(mouseWheel)
		if c.Axis == "horizontal" {
			flag = mouseHWheel
		}
		return sendMouse(flag, 0, 0, uint32(int32(c.Steps*120)))
	}
	return errors.New("unsupported desktop action")
}
func openDashboard() { _ = shellOpen("http://127.0.0.1:17993/", "") }
