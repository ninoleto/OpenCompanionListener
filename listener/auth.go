package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type replayProtector struct {
	mu   sync.Mutex
	seen map[string]time.Time
	now  func() time.Time
}

func newReplayProtector() *replayProtector {
	return &replayProtector{seen: make(map[string]time.Time), now: time.Now}
}
func macPayload(method, path, ts, nonce string, body []byte) string {
	d := sha256.Sum256(body)
	return method + "\n" + path + "\n" + ts + "\n" + nonce + "\n" + hex.EncodeToString(d[:])
}
func signature(key []byte, method, path, ts, nonce string, body []byte) string {
	m := hmac.New(sha256.New, key)
	_, _ = io.WriteString(m, macPayload(method, path, ts, nonce, body))
	return hex.EncodeToString(m.Sum(nil))
}
func (r *replayProtector) verify(req *http.Request, body []byte, keyHex string) error {
	ts := req.Header.Get("X-OCL-Timestamp")
	nonce := req.Header.Get("X-OCL-Nonce")
	sig := req.Header.Get("X-OCL-Signature")
	if len(nonce) != 32 || len(ts) < 9 || len(ts) > 12 || len(sig) != 64 {
		return errors.New("missing authentication")
	}
	if _, err := hex.DecodeString(nonce); err != nil {
		return errors.New("invalid nonce")
	}
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return errors.New("invalid configured key")
	}
	timeInt, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return errors.New("invalid timestamp")
	}
	now := r.now()
	if d := now.Sub(time.Unix(timeInt, 0)); d > 60*time.Second || d < -60*time.Second {
		return errors.New("clock skew")
	}
	expected, err := hex.DecodeString(signature(key, req.Method, req.URL.Path, ts, nonce, body))
	if err != nil {
		return err
	}
	got, err := hex.DecodeString(sig)
	if err != nil || !hmac.Equal(expected, got) {
		return errors.New("bad signature")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	for k, v := range r.seen {
		if now.Sub(v) > 2*time.Minute {
			delete(r.seen, k)
		}
	}
	if _, ok := r.seen[nonce]; ok {
		return errors.New("replay")
	}
	r.seen[nonce] = now
	return nil
}
