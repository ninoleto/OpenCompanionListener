package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	ListenIP  string `json:"listen_ip"`
	Port      int    `json:"port"`
	TrustedIP string `json:"trusted_ip"`
	Secret    string `json:"secret"`
}

func newConfig() (Config, error) {
	var b [32]byte
	if _, err := rand.Read(b[:]); err != nil {
		return Config{}, err
	}
	return Config{ListenIP: "127.0.0.1", Port: 17992, TrustedIP: "127.0.0.1", Secret: hex.EncodeToString(b[:])}, nil
}
func validateConfig(c Config) error {
	if net.ParseIP(c.ListenIP) == nil || net.ParseIP(c.TrustedIP) == nil {
		return errors.New("listen IP and trusted IP must be literal IP addresses")
	}
	if c.Port < 1024 || c.Port > 65535 {
		return errors.New("port must be between 1024 and 65535")
	}
	if _, err := hex.DecodeString(c.Secret); err != nil || len(c.Secret) != 64 {
		return errors.New("invalid pairing key")
	}
	return nil
}
func configPath() string {
	base := os.Getenv("LOCALAPPDATA")
	if base == "" {
		var err error
		base, err = os.UserConfigDir()
		if err != nil {
			base = "."
		}
	}
	return filepath.Join(base, "OpenCompanionListener", "settings.json")
}
func loadConfig(path string) (Config, error) {
	b, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		c, e := newConfig()
		if e != nil {
			return c, e
		}
		if e = saveConfig(path, c); e != nil {
			return c, e
		}
		return c, nil
	}
	if err != nil {
		return Config{}, err
	}
	var c Config
	if err = json.Unmarshal(b, &c); err != nil {
		return c, err
	}
	return c, validateConfig(c)
}
func saveConfig(path string, c Config) error {
	if err := validateConfig(c); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err = os.WriteFile(tmp, b, 0600); err != nil {
		return err
	}
	if err = os.Rename(tmp, path); err != nil {
		_ = os.Remove(tmp)
		return err
	}
	return nil
}
func addr(c Config) string { return net.JoinHostPort(c.ListenIP, strconv.Itoa(c.Port)) }
func isTrustedIP(actual string, allowed string) bool {
	// Reject ambiguous forwarded-for values. Only remote TCP peer is authoritative.
	host, _, err := net.SplitHostPort(actual)
	if err != nil {
		return false
	}
	a, b := net.ParseIP(strings.Trim(host, "[]")), net.ParseIP(allowed)
	return a != nil && b != nil && a.Equal(b)
}
