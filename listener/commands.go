package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Command struct {
	Type   string `json:"type"`
	Keys   string `json:"keys,omitempty"`
	Button string `json:"button,omitempty"`
	Steps  int    `json:"steps,omitempty"`
	Axis   string `json:"axis,omitempty"`
	DX     int    `json:"dx,omitempty"`
	DY     int    `json:"dy,omitempty"`
	URL    string `json:"url,omitempty"`
	Path   string `json:"path,omitempty"`
	Args   string `json:"args,omitempty"`
	Mode   string `json:"mode,omitempty"`
}

func validateCommand(c Command) error {
	switch c.Type {
	case "hotkey", "key_down", "key_up":
		if len(c.Keys) == 0 || len(c.Keys) > 100 {
			return errors.New("invalid keys")
		}
		for _, k := range strings.Split(c.Keys, "+") {
			if _, ok := keyCode(strings.TrimSpace(k)); !ok {
				return fmt.Errorf("unsupported key %q", k)
			}
		}
	case "num_lock":
		if c.Mode != "toggle" && c.Mode != "on" && c.Mode != "off" {
			return errors.New("Num Lock mode must be toggle, on or off")
		}
	case "click", "mouse_down", "mouse_up":
		if c.Button != "left" && c.Button != "right" && c.Button != "middle" {
			return errors.New("button must be left, right or middle")
		}
	case "scroll":
		if c.Axis != "vertical" && c.Axis != "horizontal" {
			return errors.New("axis must be vertical or horizontal")
		}
		if c.Steps == 0 || c.Steps > 20 || c.Steps < -20 {
			return errors.New("scroll steps outside -20..20")
		}
	case "move":
		if c.DX > 10000 || c.DX < -10000 || c.DY > 10000 || c.DY < -10000 {
			return errors.New("movement out of range")
		}
	case "open_url":
		u, err := url.Parse(c.URL)
		if err != nil || u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") {
			return errors.New("only absolute http/https URLs allowed")
		}
		if len(c.URL) > 2048 {
			return errors.New("URL too long")
		}
	case "open_file":
		if c.Path == "" || len(c.Path) > 1024 || strings.ContainsAny(c.Path, "\r\n\x00") || !utf8.ValidString(c.Path) {
			return errors.New("invalid file path")
		}
		if len(c.Args) > 1024 || strings.ContainsAny(c.Args, "\r\n\x00") {
			return errors.New("invalid arguments")
		}
	default:
		return fmt.Errorf("unsupported action type %q", c.Type)
	}
	return nil
}
