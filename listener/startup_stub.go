//go:build !windows

package main

import "errors"

func getStartupEnabled() (bool, error)     { return false, errors.New("Windows only") }
func setStartupEnabled(enabled bool) error { return errors.New("Windows only") }
