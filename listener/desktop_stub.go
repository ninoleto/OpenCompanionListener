//go:build !windows

package main

import (
	"errors"
	"net"
	"net/http"
	"os"
)

func executeDesktop(c Command) error {
	return errors.New("desktop input supported only on Windows")
}
func openDashboard() {}

func runApplicationLoop(a *app, ui *http.Server, ln net.Listener) error { return ui.Serve(ln) }
func requestAppExit()                                                   { os.Exit(0) }
