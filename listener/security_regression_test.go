package main

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

func securityApp(t *testing.T) *app {
	t.Helper()
	c, err := newConfig()
	if err != nil {
		t.Fatal(err)
	}
	a := &app{config: c, configFile: filepath.Join(t.TempDir(), "settings.json"), replay: newReplayProtector(), started: time.Now()}
	if err := saveConfig(a.configFile, c); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if a.server != nil {
			_ = a.server.Close()
		}
	})
	return a
}
func localPost(a *app, path string, value any) *httptest.ResponseRecorder {
	body, _ := json.Marshal(value)
	r := httptest.NewRequest("POST", "http://127.0.0.1:17993"+path, bytes.NewReader(body))
	r.Header.Set("X-OCL-Local-UI", "1")
	r.Header.Set("Origin", "http://127.0.0.1:17993")
	w := httptest.NewRecorder()
	a.uiHandler(w, r)
	return w
}
func statusCode(t *testing.T, a *app, key string) int {
	t.Helper()
	w := httptest.NewRecorder()
	a.authHandler(w, signedRequest(t, key, "GET", "/v1/status", "127.0.0.1:6000", nil))
	return w.Code
}
func TestDashboardAuthorityAndReadProtection(t *testing.T) {
	a := securityApp(t)
	for _, tc := range []struct {
		host, origin, site string
		header             bool
		want               int
	}{
		{"127.0.0.1:17993", "", "same-origin", true, 200},
		{"rebinding.example:17993", "", "same-origin", true, 403},
		{"localhost:17993", "", "same-origin", true, 403},
		{"127.0.0.1:80", "", "same-origin", true, 403},
		{"", "", "same-origin", true, 403},
		{"127.0.0.1:17993", "http://foreign.example", "", true, 403},
		{"127.0.0.1:17993", "null", "", true, 403},
		{"127.0.0.1:17993", "", "cross-site", true, 403},
		{"127.0.0.1:17993", "", "", false, 403},
	} {
		r := httptest.NewRequest("GET", "/ui/state", nil)
		r.Host = tc.host
		r.Header.Set("Origin", tc.origin)
		r.Header.Set("Sec-Fetch-Site", tc.site)
		if tc.header {
			r.Header.Set("X-OCL-Local-UI", "1")
		}
		w := httptest.NewRecorder()
		a.uiHandler(w, r)
		if w.Code != tc.want {
			t.Errorf("authority %q origin %q header %v: got %d want %d", tc.host, tc.origin, tc.header, w.Code, tc.want)
		}
		if tc.want != 200 && bytes.Contains(w.Body.Bytes(), []byte(a.config.Secret)) {
			t.Error("pairing key disclosed; redacted")
		}
	}
	// Reject the foreign authority before any route can perform a mutation.
	old := a.config
	r := httptest.NewRequest("POST", "http://foreign.example:17993/ui/rotate", nil)
	r.Header.Set("X-OCL-Local-UI", "1")
	w := httptest.NewRecorder()
	a.uiHandler(w, r)
	if w.Code != 403 || a.config != old {
		t.Fatal("foreign authority mutated settings")
	}
	for _, path := range []string{"/", "/icon.png"} {
		r = httptest.NewRequest("GET", "http://127.0.0.1:17993"+path, nil)
		w = httptest.NewRecorder()
		a.uiHandler(w, r)
		if w.Code != 200 {
			t.Fatalf("dashboard asset %s unavailable", path)
		}
	}
}
func TestSettingsSavePreservesReplay(t *testing.T) {
	a := securityApp(t)
	r := signedRequest(t, a.config.Secret, "GET", "/v1/status", "127.0.0.1:6000", nil)
	w := httptest.NewRecorder()
	a.authHandler(w, r)
	if w.Code != 200 {
		t.Fatal("initial status failed")
	}
	if w = localPost(a, "/ui/save", a.config); w.Code != 200 {
		t.Fatalf("save: %d", w.Code)
	}
	w = httptest.NewRecorder()
	a.authHandler(w, r)
	if w.Code != 401 {
		t.Fatalf("replay after save accepted: %d", w.Code)
	}
}
func TestRestartPreservesReplay(t *testing.T) {
	a := securityApp(t)
	c := a.config
	c.Port = 0 // Bind an ephemeral test port; never use the live API port.
	if err := a.start(c); err != nil {
		t.Fatal(err)
	}
	r := signedRequest(t, c.Secret, "GET", "/v1/status", "127.0.0.1:6000", nil)
	w := httptest.NewRecorder()
	a.authHandler(w, r)
	if w.Code != 200 {
		t.Fatal("initial status failed")
	}
	if err := a.restart(); err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	a.authHandler(w, r)
	if w.Code != 401 {
		t.Fatalf("replay after restart accepted: %d", w.Code)
	}
}
func TestRotationPersistenceFailureKeepsOldKey(t *testing.T) {
	a := securityApp(t)
	old := a.config
	originalPath := a.configFile
	// A nonempty directory cannot be replaced by a settings file.
	a.configFile = filepath.Join(t.TempDir(), "blocked")
	if err := os.Mkdir(a.configFile, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(a.configFile, "keep"), []byte("fixture"), 0600); err != nil {
		t.Fatal(err)
	}
	w := localPost(a, "/ui/rotate", nil)
	if w.Code != 500 {
		t.Fatalf("expected failed write: %d", w.Code)
	}
	if a.config != old || statusCode(t, a, old.Secret) != 200 {
		t.Fatal("failed rotation invalidated active key; redacted")
	}
	disk, err := loadConfig(originalPath)
	if err != nil || disk != old {
		t.Fatal("persisted configuration changed")
	}
}
func TestSuccessfulRotationPersistsAndRevokesOldKey(t *testing.T) {
	a := securityApp(t)
	old := a.config.Secret
	if w := localPost(a, "/ui/rotate", nil); w.Code != 200 {
		t.Fatalf("rotation: %d", w.Code)
	}
	disk, err := loadConfig(a.configFile)
	if err != nil || disk != a.config || disk.Secret == old {
		t.Fatal("rotation not persisted; redacted")
	}
	if statusCode(t, a, old) != 401 || statusCode(t, a, disk.Secret) != 200 {
		t.Fatal("rotation authentication incorrect")
	}
}
func TestFailedNetworkSaveLeavesListenerAndConfig(t *testing.T) {
	a := securityApp(t)
	c := a.config
	c.Port = 0
	if err := a.start(c); err != nil {
		t.Fatal(err)
	}
	// Record the real ephemeral port so the old config remains valid.
	a.config.Port = a.listener.Addr().(*net.TCPAddr).Port
	if err := saveConfig(a.configFile, a.config); err != nil {
		t.Fatal(err)
	}
	old := a.config
	oldListener := a.listener
	candidate := old
	candidate.TrustedIP = "192.0.2.10"
	a.configFile = t.TempDir()
	if w := localPost(a, "/ui/save", candidate); w.Code != 500 {
		t.Fatalf("expected failure: %d", w.Code)
	}
	if a.config != old || a.listener != oldListener {
		t.Fatal("failed network save changed active configuration")
	}
}
func TestConcurrentSaveAndRotateRemainConsistent(t *testing.T) {
	a := securityApp(t)
	network := a.config
	var wg sync.WaitGroup
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			route := "/ui/save"
			if i%2 == 0 {
				route = "/ui/rotate"
			}
			if w := localPost(a, route, network); w.Code != 200 {
				t.Errorf("concurrent %s failed: %d", route, w.Code)
			}
		}(i)
	}
	wg.Wait()
	disk, err := loadConfig(a.configFile)
	if err != nil || disk != a.config {
		t.Fatal("disk and active configuration diverged; redacted")
	}
}

func TestFailedEndpointChangeReleasesPreparedPort(t *testing.T) {
	a := securityApp(t)
	c := a.config
	c.Port = 0
	if err := a.start(c); err != nil {
		t.Fatal(err)
	}
	a.config.Port = a.listener.Addr().(*net.TCPAddr).Port
	old, oldListener := a.config, a.listener
	probe, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	candidate := old
	candidate.Port = probe.Addr().(*net.TCPAddr).Port
	// A bind failure must leave the old listener and persisted settings alone.
	if w := localPost(a, "/ui/save", candidate); w.Code != 409 {
		_ = probe.Close()
		t.Fatalf("expected bind conflict: %d", w.Code)
	}
	_ = probe.Close()
	a.configFile = t.TempDir() // Rename over a directory must fail.
	if w := localPost(a, "/ui/save", candidate); w.Code != 500 {
		t.Fatalf("expected persistence failure: %d", w.Code)
	}
	if a.config != old || a.listener != oldListener {
		t.Fatal("failed endpoint change disturbed the active listener")
	}
	probe, err = net.Listen("tcp", addr(candidate))
	if err != nil {
		t.Fatalf("prepared port leaked after persistence failure: %v", err)
	}
	_ = probe.Close()
	if statusCode(t, a, old.Secret) != 200 {
		t.Fatal("old authentication stopped working")
	}
}
