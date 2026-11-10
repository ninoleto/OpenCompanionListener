package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func signedRequest(t *testing.T, key, method, path, remote string, body []byte) *http.Request {
	t.Helper()
	req := httptest.NewRequest(method, path, bytes.NewReader(body))
	req.RemoteAddr = remote
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		t.Fatal(err)
	}
	nonce := hex.EncodeToString(b)
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	req.Header.Set("X-OCL-Nonce", nonce)
	req.Header.Set("X-OCL-Timestamp", ts)
	req.Header.Set("X-OCL-Signature", signature(mustDecode(t, key), method, path, ts, nonce, body))
	return req
}
func mustDecode(t *testing.T, s string) []byte {
	t.Helper()
	b, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
func TestAuthenticationAndReplay(t *testing.T) {
	c, err := newConfig()
	if err != nil {
		t.Fatal(err)
	}
	c.TrustedIP = "192.0.2.17"
	a := &app{config: c, replay: newReplayProtector()}
	req := signedRequest(t, c.Secret, "GET", "/v1/status", "192.0.2.17:6000", nil)
	w := httptest.NewRecorder()
	a.authHandler(w, req)
	if w.Code != 200 {
		t.Fatalf("expected 200; %d %s", w.Code, w.Body.String())
	}
	w = httptest.NewRecorder()
	a.authHandler(w, req)
	if w.Code != 401 {
		t.Fatalf("replayed request accepted: %d", w.Code)
	}
	bad := signedRequest(t, c.Secret, "GET", "/v1/status", "192.0.2.99:6000", nil)
	w = httptest.NewRecorder()
	a.authHandler(w, bad)
	if w.Code != 403 {
		t.Fatalf("other IP accepted: %d", w.Code)
	}
}
func TestChangedPayloadAndStaleTimeRejected(t *testing.T) {
	c, _ := newConfig()
	r := newReplayProtector()
	req := signedRequest(t, c.Secret, "POST", "/v1/command", "127.0.0.1:6000", []byte(`{"type":"scroll","axis":"vertical","steps":1}`))
	if err := r.verify(req, []byte(`{"type":"scroll","axis":"vertical","steps":10}`), c.Secret); err == nil {
		t.Fatal("modified body accepted")
	}
	req = signedRequest(t, c.Secret, "GET", "/v1/status", "127.0.0.1:6000", nil)
	req.Header.Set("X-OCL-Timestamp", strconv.FormatInt(time.Now().Unix()-200, 10))
	if err := r.verify(req, nil, c.Secret); err == nil {
		t.Fatal("stale request accepted")
	}
}
func TestValidation(t *testing.T) {
	valid := []Command{{Type: "hotkey", Keys: "CTRL+SHIFT+S"}, {Type: "click", Button: "left"}, {Type: "scroll", Axis: "vertical", Steps: -3}, {Type: "open_url", URL: "https://example.com/"}, {Type: "open_file", Path: `C:\Windows\notepad.exe`}}
	for _, c := range valid {
		if err := validateCommand(c); err != nil {
			t.Errorf("%+v invalid: %v", c, err)
		}
	}
	invalid := []Command{{Type: "shell"}, {Type: "open_url", URL: "file:///C:/Windows/system.ini"}, {Type: "open_url", URL: "javascript:alert(1)"}, {Type: "hotkey", Keys: "CTRL+BANANA"}, {Type: "scroll", Steps: 200, Axis: "vertical"}, {Type: "open_file", Path: "bad\npath"}}
	for _, c := range invalid {
		if err := validateCommand(c); err == nil {
			t.Errorf("%+v unexpectedly accepted", c)
		}
	}
}
func TestConfigRoundTrip(t *testing.T) {
	c, _ := newConfig()
	path := t.TempDir() + "/settings.json"
	if err := saveConfig(path, c); err != nil {
		t.Fatal(err)
	}
	loaded, err := loadConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	if loaded != c {
		t.Fatalf("got %+v, want %+v", loaded, c)
	}
}
