//go:build windows

package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// Per-user Run key. Does not need elevation or a Windows service.
const runKeyPath = `Software\Microsoft\Windows\CurrentVersion\Run`
const runValueName = "OpenCompanionListener"
const (
	hkeyCurrentUser   = uintptr(0x80000001)
	keyQueryValue     = uintptr(0x0001)
	keySetValue       = uintptr(0x0002)
	regSZ             = uintptr(1)
	errorFileNotFound = syscall.Errno(2)
)

var startupAdvapi = syscall.NewLazyDLL("advapi32.dll")
var startupOpenKey = startupAdvapi.NewProc("RegOpenKeyExW")
var startupCloseKey = startupAdvapi.NewProc("RegCloseKey")
var startupSetValue = startupAdvapi.NewProc("RegSetValueExW")
var startupQueryValue = startupAdvapi.NewProc("RegQueryValueExW")
var startupDeleteValue = startupAdvapi.NewProc("RegDeleteValueW")

func openRunKey(access uintptr) (uintptr, error) {
	path, _ := syscall.UTF16PtrFromString(runKeyPath)
	var handle uintptr
	code, _, _ := startupOpenKey.Call(hkeyCurrentUser, uintptr(unsafe.Pointer(path)), 0, access, uintptr(unsafe.Pointer(&handle)))
	if code != 0 {
		return 0, syscall.Errno(code)
	}
	return handle, nil
}
func startupCommand() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return `"` + strings.ReplaceAll(exe, `"`, "") + `"`, nil
}
func getStartupEnabled() (bool, error) {
	handle, err := openRunKey(keyQueryValue)
	if err != nil {
		return false, err
	}
	defer startupCloseKey.Call(handle)
	name, _ := syscall.UTF16PtrFromString(runValueName)
	var typ uint32
	var length uint32
	code, _, _ := startupQueryValue.Call(handle, uintptr(unsafe.Pointer(name)), 0, uintptr(unsafe.Pointer(&typ)), 0, uintptr(unsafe.Pointer(&length)))
	if syscall.Errno(code) == errorFileNotFound {
		return false, nil
	}
	if code != 0 {
		return false, syscall.Errno(code)
	}
	if typ != uint32(regSZ) || length == 0 || length > 32768 {
		return false, fmt.Errorf("unexpected Run value type or size")
	}
	buf := make([]uint16, (int(length)+1)/2)
	code, _, _ = startupQueryValue.Call(handle, uintptr(unsafe.Pointer(name)), 0, uintptr(unsafe.Pointer(&typ)), uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&length)))
	if code != 0 {
		return false, syscall.Errno(code)
	}
	want, err := startupCommand()
	if err != nil {
		return false, err
	}
	return strings.EqualFold(strings.TrimSpace(syscall.UTF16ToString(buf)), want), nil
}
func setStartupEnabled(enabled bool) error {
	handle, err := openRunKey(keySetValue)
	if err != nil {
		return err
	}
	defer startupCloseKey.Call(handle)
	name, _ := syscall.UTF16PtrFromString(runValueName)
	if !enabled {
		code, _, _ := startupDeleteValue.Call(handle, uintptr(unsafe.Pointer(name)))
		if code != 0 && syscall.Errno(code) != errorFileNotFound {
			return syscall.Errno(code)
		}
		return nil
	}
	cmd, err := startupCommand()
	if err != nil {
		return err
	}
	value, err := syscall.UTF16FromString(cmd)
	if err != nil {
		return err
	}
	code, _, _ := startupSetValue.Call(handle, uintptr(unsafe.Pointer(name)), 0, regSZ, uintptr(unsafe.Pointer(&value[0])), uintptr(len(value)*2))
	if code != 0 {
		return syscall.Errno(code)
	}
	return nil
}
