//go:build windows

package main

// Native Windows notification-area UI. No browser runtime, framework, or
// shell subprocess is required for the tray itself.
import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"unsafe"
)

const (
	trayMessage       = 0x8001 // WM_APP + 1
	wmClose           = 0x0010
	wmDestroy         = 0x0002
	wmLButtonDblClk   = 0x0203
	wmRButtonUp       = 0x0205
	wmContextMenu     = 0x007B
	ninSelect         = 0x0400
	ninKeySelect      = 0x0401
	nimAdd            = 0
	nimDelete         = 2
	nimSetVersion     = 4
	nifMessage        = 1
	nifIcon           = 2
	nifTip            = 4
	menuString        = 0
	menuSeparator     = 0x800
	menuReturnCommand = 0x100
	menuRightButton   = 2
	trayMenuOpen      = 201
	trayMenuRestart   = 202
	trayMenuExit      = 203
)

type winPoint struct{ X, Y int32 }
type winMessage struct {
	Hwnd    uintptr
	Msg     uint32
	_       uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Point   winPoint
	Private uint32
}
type wndClassEx struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   uintptr
	Icon       uintptr
	Cursor     uintptr
	Background uintptr
	MenuName   *uint16
	ClassName  *uint16
	IconSmall  uintptr
}
type notifyIconData struct {
	Size            uint32
	_               uint32
	Hwnd            uintptr
	ID              uint32
	Flags           uint32
	CallbackMessage uint32
	_               uint32
	Icon            uintptr
	Tip             [128]uint16
	State           uint32
	StateMask       uint32
	Info            [256]uint16
	Version         uint32
	InfoTitle       [64]uint16
	InfoFlags       uint32
	GUID            [16]byte
	BalloonIcon     uintptr
}

var (
	trayUser32            = syscall.NewLazyDLL("user32.dll")
	trayShell32           = syscall.NewLazyDLL("shell32.dll")
	trayKernel32          = syscall.NewLazyDLL("kernel32.dll")
	registerClassEx       = trayUser32.NewProc("RegisterClassExW")
	createWindowEx        = trayUser32.NewProc("CreateWindowExW")
	defWindowProc         = trayUser32.NewProc("DefWindowProcW")
	destroyWindow         = trayUser32.NewProc("DestroyWindow")
	getMessage            = trayUser32.NewProc("GetMessageW")
	translateMessage      = trayUser32.NewProc("TranslateMessage")
	dispatchMessage       = trayUser32.NewProc("DispatchMessageW")
	postQuitMessage       = trayUser32.NewProc("PostQuitMessage")
	postMessage           = trayUser32.NewProc("PostMessageW")
	getCursorPos          = trayUser32.NewProc("GetCursorPos")
	setForegroundWindow   = trayUser32.NewProc("SetForegroundWindow")
	createPopupMenu       = trayUser32.NewProc("CreatePopupMenu")
	appendMenu            = trayUser32.NewProc("AppendMenuW")
	trackPopupMenu        = trayUser32.NewProc("TrackPopupMenu")
	destroyMenu           = trayUser32.NewProc("DestroyMenu")
	loadImage             = trayUser32.NewProc("LoadImageW")
	loadIcon              = trayUser32.NewProc("LoadIconW")
	destroyIcon           = trayUser32.NewProc("DestroyIcon")
	registerWindowMessage = trayUser32.NewProc("RegisterWindowMessageW")
	getModuleHandle       = trayKernel32.NewProc("GetModuleHandleW")
	notifyIcon            = trayShell32.NewProc("Shell_NotifyIconW")
	activeTray            *app
	activeTrayHandle      uintptr
	activeTrayIcon        uintptr
	activeTrayIconOwned   bool
	taskbarCreatedMessage uint32
)

func trayData(hwnd uintptr, icon uintptr) notifyIconData {
	data := notifyIconData{Size: uint32(unsafe.Sizeof(notifyIconData{})), Hwnd: hwnd, ID: 1,
		Flags: nifMessage | nifIcon | nifTip, CallbackMessage: trayMessage, Icon: icon}
	copy(data.Tip[:], syscall.StringToUTF16("Open Companion Listener"))
	return data
}
func trayNotify(operation uintptr, data *notifyIconData) bool {
	ok, _, _ := notifyIcon.Call(operation, uintptr(unsafe.Pointer(data)))
	return ok != 0
}
func addTrayIcon() bool {
	data := trayData(activeTrayHandle, activeTrayIcon)
	if !trayNotify(nimAdd, &data) {
		return false
	}
	data.Version = 4
	_ = trayNotify(nimSetVersion, &data)
	return true
}
func showTrayMenu(hwnd uintptr) {
	menu, _, _ := createPopupMenu.Call()
	if menu == 0 {
		return
	}
	defer destroyMenu.Call(menu)
	add := func(flags uintptr, id uintptr, label string) {
		p, _ := syscall.UTF16PtrFromString(label)
		_, _, _ = appendMenu.Call(menu, flags, id, uintptr(unsafe.Pointer(p)))
	}
	add(menuString, trayMenuOpen, "Open GUI")
	add(menuString, trayMenuRestart, "Restart Listener")
	add(menuSeparator, 0, "")
	add(menuString, trayMenuExit, "Exit")
	var pt winPoint
	ok, _, _ := getCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if ok == 0 {
		return
	}
	setForegroundWindow.Call(hwnd)
	cmd, _, _ := trackPopupMenu.Call(menu, menuReturnCommand|menuRightButton,
		uintptr(pt.X), uintptr(pt.Y), 0, hwnd, 0)
	switch cmd {
	case trayMenuOpen:
		openDashboard()
	case trayMenuRestart:
		if activeTray != nil {
			_ = activeTray.restart()
		}
	case trayMenuExit:
		destroyWindow.Call(hwnd)
	}
	postMessage.Call(hwnd, 0, 0, 0)
}
func trayWindowProc(hwnd uintptr, message uint32, wparam, lparam uintptr) uintptr {
	if message == taskbarCreatedMessage && taskbarCreatedMessage != 0 {
		addTrayIcon()
		return 0
	}
	switch message {
	case trayMessage:
		action := uint32(lparam & 0xffff) // NIM_SETVERSION v4
		switch action {
		case wmLButtonDblClk, ninSelect, ninKeySelect:
			openDashboard()
		case wmRButtonUp, wmContextMenu:
			showTrayMenu(hwnd)
		}
		return 0
	case wmClose:
		destroyWindow.Call(hwnd)
		return 0
	case wmDestroy:
		postQuitMessage.Call(0)
		return 0
	}
	result, _, _ := defWindowProc.Call(hwnd, uintptr(message), wparam, lparam)
	return result
}
func loadTrayIcon() (uintptr, bool) {
	instance, _, _ := getModuleHandle.Call(0)
	embedded, _, _ := loadIcon.Call(instance, 1) // RT_GROUP_ICON ID 1, embedded in EXE
	if embedded != 0 {
		return embedded, false
	}
	exe, err := os.Executable()
	if err == nil {
		iconPath := filepath.Join(filepath.Dir(exe), "OpenCompanionListener.ico")
		p, err := syscall.UTF16PtrFromString(iconPath)
		if err == nil {
			h, _, _ := loadImage.Call(0, uintptr(unsafe.Pointer(p)), 1, 0, 0, 0x10|0x40)
			if h != 0 {
				return h, true
			}
		}
	}
	h, _, _ := loadIcon.Call(0, 32512) // Windows stock icon if user moves EXE alone
	return h, false
}
func runTray(a *app) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	activeTray = a
	defer func() { activeTray = nil; activeTrayHandle = 0 }()
	classname, _ := syscall.UTF16PtrFromString("OpenCompanionListenerTray_012")
	title, _ := syscall.UTF16PtrFromString("Open Companion Listener")
	taskname, _ := syscall.UTF16PtrFromString("TaskbarCreated")
	created, _, _ := registerWindowMessage.Call(uintptr(unsafe.Pointer(taskname)))
	taskbarCreatedMessage = uint32(created)
	instance, _, _ := getModuleHandle.Call(0)
	icon, owned := loadTrayIcon()
	activeTrayIcon, activeTrayIconOwned = icon, owned
	if owned {
		defer destroyIcon.Call(icon)
	}
	wc := wndClassEx{Size: uint32(unsafe.Sizeof(wndClassEx{})), WndProc: syscall.NewCallback(trayWindowProc),
		Instance: instance, ClassName: classname, Icon: icon, IconSmall: icon}
	cls, _, err := registerClassEx.Call(uintptr(unsafe.Pointer(&wc)))
	if cls == 0 {
		return fmt.Errorf("register tray window: %w", err)
	}
	hwnd, _, err := createWindowEx.Call(0, uintptr(unsafe.Pointer(classname)), uintptr(unsafe.Pointer(title)), 0,
		0, 0, 0, 0, 0, 0, instance, 0)
	if hwnd == 0 {
		return fmt.Errorf("create tray window: %w", err)
	}
	activeTrayHandle = hwnd
	if !addTrayIcon() {
		destroyWindow.Call(hwnd)
		return errors.New("cannot add Windows tray icon")
	}
	defer func() { data := trayData(hwnd, 0); trayNotify(nimDelete, &data) }()
	var msg winMessage
	for {
		r, _, e := getMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		if int32(r) == -1 {
			return fmt.Errorf("message loop: %w", e)
		}
		if r == 0 {
			return nil
		}
		translateMessage.Call(uintptr(unsafe.Pointer(&msg)))
		dispatchMessage.Call(uintptr(unsafe.Pointer(&msg)))
	}
}
func runApplicationLoop(a *app, ui *http.Server, ln net.Listener) error {
	failed := make(chan error, 1)
	go func() { failed <- ui.Serve(ln) }()
	if err := runTray(a); err != nil {
		a.appendEvent("error", "Tray unavailable: "+err.Error())
		// Do not terminate a working listener solely because Explorer does
		// not accept a notification-area icon.
		return <-failed
	}
	_ = ui.Close()
	a.mu.RLock()
	server := a.server
	apiLn := a.listener
	a.mu.RUnlock()
	if apiLn != nil {
		_ = apiLn.Close()
	}
	if server != nil {
		_ = server.Close()
	}
	return nil
}
func requestAppExit() {
	hwnd := activeTrayHandle
	if hwnd != 0 {
		postMessage.Call(hwnd, wmClose, 0, 0)
	} else {
		os.Exit(0)
	}
}
