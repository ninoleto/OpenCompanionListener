package main

import (
	"bytes"
	"crypto/rand"
	"embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

//go:embed ui/index.html ui/icon.png
var uiFiles embed.FS

type event struct {
	At      string `json:"at"`
	Message string `json:"message"`
	Kind    string `json:"kind"`
}
type app struct {
	// Serialize persistence and listener lifecycle changes. Never hold mu while
	// acquiring settingsMu; API requests only need mu for their state snapshot.
	settingsMu sync.Mutex
	mu         sync.RWMutex
	config     Config
	configFile string
	server     *http.Server
	listener   net.Listener
	lastError  string
	events     []event
	replay     *replayProtector
	started    time.Time
}

func (a *app) appendEvent(kind, message string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.events = append(a.events, event{At: time.Now().Format("15:04:05"), Message: message, Kind: kind})
	if len(a.events) > 80 {
		a.events = append([]event(nil), a.events[len(a.events)-80:]...)
	}
}
func (a *app) start(c Config) error {
	a.settingsMu.Lock()
	defer a.settingsMu.Unlock()
	return a.startLocked(c)
}
func (a *app) startLocked(c Config) error {
	ln, err := net.Listen("tcp", addr(c))
	if err != nil {
		return err
	}
	a.installListener(c, ln)
	return nil
}

// installListener publishes a prepared listener after all fallible preparation
// (including settings persistence, when needed) has succeeded.
func (a *app) installListener(c Config, ln net.Listener) {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/status", a.authHandler)
	mux.HandleFunc("/v1/command", a.authHandler)
	s := &http.Server{Handler: mux, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 30 * time.Second, MaxHeaderBytes: 8192}
	a.mu.Lock()
	old := a.server
	oldLn := a.listener
	a.server = s
	a.listener = ln
	a.config = c
	a.lastError = ""
	a.started = time.Now()
	if a.replay == nil {
		a.replay = newReplayProtector()
	}
	a.mu.Unlock()
	if oldLn != nil {
		_ = oldLn.Close()
	}
	if old != nil {
		_ = old.Close()
	}
	go func() {
		if err := s.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) && !errors.Is(err, net.ErrClosed) {
			a.appendEvent("error", err.Error())
		}
	}()
	a.appendEvent("ok", "Listening on "+addr(c)+" · trusted "+c.TrustedIP)
}

// restart closes the API listener before rebinding the same address. The
// local dashboard and tray remain running even if the API rebinding fails.
func (a *app) restart() error {
	a.settingsMu.Lock()
	defer a.settingsMu.Unlock()
	a.mu.Lock()
	c := a.config
	old := a.server
	oldLn := a.listener
	a.server, a.listener = nil, nil
	a.mu.Unlock()
	if oldLn != nil {
		_ = oldLn.Close()
	}
	if old != nil {
		_ = old.Close()
	}
	if err := a.startLocked(c); err != nil {
		a.mu.Lock()
		a.lastError = err.Error()
		a.mu.Unlock()
		a.appendEvent("error", "Listener restart failed: "+err.Error())
		return err
	}
	a.appendEvent("ok", "Listener restarted")
	return nil
}
func (a *app) authHandler(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path == "/v1/command" && r.Method != "POST") || (r.URL.Path == "/v1/status" && r.Method != "GET") {
		http.Error(w, "method", http.StatusMethodNotAllowed)
		return
	}
	a.mu.RLock()
	c := a.config
	replay := a.replay
	a.mu.RUnlock()
	if !isTrustedIP(r.RemoteAddr, c.TrustedIP) {
		http.Error(w, "not trusted", http.StatusForbidden)
		return
	}
	body, err := io.ReadAll(io.LimitReader(r.Body, 8193))
	if err != nil || len(body) > 8192 {
		http.Error(w, "body too large", http.StatusRequestEntityTooLarge)
		return
	}
	if err = replay.verify(r, body, c.Secret); err != nil {
		http.Error(w, "authentication failed", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	if r.URL.Path == "/v1/status" {
		_, _ = io.WriteString(w, `{"ok":true,"version":"0.1.4","platform":"windows"}`)
		return
	}
	var cmd Command
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	if err = dec.Decode(&cmd); err == nil {
		var extra any
		if dec.Decode(&extra) != io.EOF {
			err = errors.New("extra JSON content")
		}
	}
	if err == nil {
		err = validateCommand(cmd)
	}
	if err != nil {
		http.Error(w, "invalid command: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err = executeDesktop(cmd); err != nil {
		a.appendEvent("error", "Action "+cmd.Type+" failed: "+err.Error())
		http.Error(w, "execution failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	a.appendEvent("action", "OS accepted "+cmd.Type+" from "+c.TrustedIP)
	_, _ = io.WriteString(w, `{"ok":true}`)
}
func (a *app) uiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Security-Policy", "default-src 'none'; img-src 'self'; script-src 'unsafe-inline'; style-src 'unsafe-inline'; connect-src 'self'; frame-ancestors 'none'; base-uri 'none'")
	// Loopback binding alone does not prevent DNS rebinding. Use the canonical
	// dashboard authority, and guard reads as well as mutations.
	const authority = "127.0.0.1:17993"
	if r.Host != authority || (r.URL.Host != "" && r.URL.Host != authority) {
		http.Error(w, "invalid dashboard host", http.StatusForbidden)
		return
	}
	if origin := r.Header.Get("Origin"); origin != "" && origin != "http://"+authority {
		http.Error(w, "invalid origin", http.StatusForbidden)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/ui/") &&
		(r.Header.Get("X-OCL-Local-UI") != "1" || r.Header.Get("Sec-Fetch-Site") == "cross-site") {
		http.Error(w, "local UI request required", http.StatusForbidden)
		return
	}
	if r.URL.Path == "/icon.png" && r.Method == "GET" {
		b, _ := uiFiles.ReadFile("ui/icon.png")
		w.Header().Set("Content-Type", "image/png")
		_, _ = w.Write(b)
		return
	}
	if r.URL.Path == "/" && r.Method == "GET" {
		b, _ := uiFiles.ReadFile("ui/index.html")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write(b)
		return
	}
	if r.URL.Path == "/ui/state" && r.Method == "GET" {
		a.mu.RLock()
		state := map[string]any{"config": a.config, "running": a.listener != nil, "error": a.lastError, "events": a.events, "uptime_seconds": int(time.Since(a.started).Seconds())}
		a.mu.RUnlock()
		startupEnabled, startupErr := getStartupEnabled()
		state["startup_enabled"] = startupEnabled
		if startupErr != nil {
			state["startup_error"] = startupErr.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(state)
		return
	}
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}
	if r.Header.Get("X-OCL-Local-UI") != "1" {
		http.Error(w, "UI header required", http.StatusForbidden)
		return
	}
	if r.URL.Path == "/ui/save" {
		var candidate struct {
			ListenIP  string `json:"listen_ip"`
			Port      int    `json:"port"`
			TrustedIP string `json:"trusted_ip"`
		}
		if err := json.NewDecoder(io.LimitReader(r.Body, 2048)).Decode(&candidate); err != nil {
			http.Error(w, "invalid settings", 400)
			return
		}
		a.settingsMu.Lock()
		defer a.settingsMu.Unlock()
		a.mu.RLock()
		old := a.config
		a.mu.RUnlock()
		c := old
		c.ListenIP = strings.TrimSpace(candidate.ListenIP)
		c.Port = candidate.Port
		c.TrustedIP = strings.TrimSpace(candidate.TrustedIP)
		if err := validateConfig(c); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		var prepared net.Listener
		if addr(c) != addr(old) {
			// Avoid losing a working server due to a typo. Bind the new endpoint first.
			var err error
			prepared, err = net.Listen("tcp", addr(c))
			if err != nil {
				http.Error(w, "Could not bind new address: "+err.Error(), 409)
				return
			}
		}
		if err := saveConfig(a.configFile, c); err != nil {
			if prepared != nil {
				_ = prepared.Close()
			}
			http.Error(w, "Could not persist settings: "+err.Error(), 500)
			return
		}
		if prepared != nil {
			a.installListener(c, prepared)
		} else {
			a.mu.Lock()
			a.config = c
			a.mu.Unlock()
		}
		a.appendEvent("ok", "Network settings saved")
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"ok":true}`)
		return
	}
	if r.URL.Path == "/ui/startup" {
		var requested struct {
			Enabled bool `json:"enabled"`
		}
		if err := json.NewDecoder(io.LimitReader(r.Body, 512)).Decode(&requested); err != nil {
			http.Error(w, "invalid startup settings", http.StatusBadRequest)
			return
		}
		if err := setStartupEnabled(requested.Enabled); err != nil {
			http.Error(w, "Windows startup update failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		a.appendEvent("ok", "Start with Windows "+map[bool]string{true: "enabled", false: "disabled"}[requested.Enabled])
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"ok":true}`)
		return
	}
	if r.URL.Path == "/ui/rotate" {
		a.settingsMu.Lock()
		defer a.settingsMu.Unlock()
		var b [32]byte
		if _, err := rand.Read(b[:]); err != nil {
			http.Error(w, "secure random unavailable", 500)
			return
		}
		a.mu.RLock()
		c := a.config
		a.mu.RUnlock()
		c.Secret = hex.EncodeToString(b[:])
		if err := saveConfig(a.configFile, c); err != nil {
			http.Error(w, "save failed", 500)
			return
		}
		a.mu.Lock()
		a.config = c
		a.mu.Unlock()
		a.appendEvent("ok", "Pairing key rotated; update Companion module")
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"ok":true}`)
		return
	}
	if r.URL.Path == "/ui/restart" {
		if err := a.restart(); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"ok":true}`)
		return
	}
	if r.URL.Path == "/ui/exit" {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"ok":true}`)
		go func() { time.Sleep(300 * time.Millisecond); requestAppExit() }()
		return
	}
	http.NotFound(w, r)
}
func main() {
	path := configPath()
	c, err := loadConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	a := &app{config: c, configFile: path, replay: newReplayProtector(), started: time.Now()}
	a.appendEvent("info", "Ready. Free and open-source desktop control.")
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.uiHandler)
	uiServer := &http.Server{Handler: mux, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second}
	uiLn, err := net.Listen("tcp", "127.0.0.1:17993")
	if err != nil {
		return
	} // Already running: do not launch a second instance or open the dashboard.
	if err = a.start(c); err != nil {
		a.mu.Lock()
		a.lastError = err.Error()
		a.mu.Unlock()
		a.appendEvent("error", "Failed to start: "+err.Error())
	}
	if err = runApplicationLoop(a, uiServer, uiLn); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("application stopped: %v", err)
	}
}
