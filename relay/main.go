// pagerr-relay: a tiny in-memory mailbox used by Pagerr's "Quick connect".
//
// It never sees a key or readable data. A receiving device opens an empty slot
// under a random ID, a sending device drops one encrypted blob into it, and the
// receiver collects it once. Slots live in memory only, expire after a few
// minutes and are deleted as soon as they are collected.
package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"
)

const (
	slotTTL      = 3 * time.Minute
	maxBlobBytes = 512 * 1024
	maxSlots     = 50
	createsPerIP = 12 // per minute
)

var idRe = regexp.MustCompile(`^[A-Za-z0-9_-]{22}$`)

type slot struct {
	data    []byte
	expires time.Time
}

type relay struct {
	mu      sync.Mutex
	slots   map[string]*slot
	creates map[string][]time.Time
}

func newRelay() *relay {
	return &relay{slots: map[string]*slot{}, creates: map[string][]time.Time{}}
}

func (r *relay) sweep() {
	now := time.Now()
	r.mu.Lock()
	defer r.mu.Unlock()
	for id, s := range r.slots {
		if now.After(s.expires) {
			delete(r.slots, id)
		}
	}
	for ip, ts := range r.creates {
		keep := ts[:0]
		for _, t := range ts {
			if now.Sub(t) < time.Minute {
				keep = append(keep, t)
			}
		}
		if len(keep) == 0 {
			delete(r.creates, ip)
		} else {
			r.creates[ip] = keep
		}
	}
}

func clientIP(req *http.Request) string {
	if v := req.Header.Get("X-Real-IP"); v != "" {
		return v
	}
	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return req.RemoteAddr
	}
	return host
}

func (r *relay) handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	if req.URL.Path == "/ping" && req.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "pagerr-relay")
		return
	}

	if !strings.HasPrefix(req.URL.Path, "/pair/") {
		http.NotFound(w, req)
		return
	}
	id := strings.TrimPrefix(req.URL.Path, "/pair/")
	if !idRe.MatchString(id) {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodPost: // receiver opens an empty slot
		ip := clientIP(req)
		r.mu.Lock()
		defer r.mu.Unlock()
		now := time.Now()
		if len(r.slots) >= maxSlots {
			http.Error(w, "busy", http.StatusTooManyRequests)
			return
		}
		ts := r.creates[ip]
		if len(ts) >= createsPerIP {
			http.Error(w, "slow down", http.StatusTooManyRequests)
			return
		}
		if _, exists := r.slots[id]; exists {
			http.Error(w, "exists", http.StatusConflict)
			return
		}
		r.creates[ip] = append(ts, now)
		r.slots[id] = &slot{expires: now.Add(slotTTL)}
		w.WriteHeader(http.StatusCreated)

	case http.MethodPut: // sender fills it, once
		body, err := io.ReadAll(http.MaxBytesReader(w, req.Body, maxBlobBytes))
		if err != nil || len(body) == 0 {
			http.Error(w, "bad body", http.StatusBadRequest)
			return
		}
		r.mu.Lock()
		defer r.mu.Unlock()
		s, ok := r.slots[id]
		if !ok || time.Now().After(s.expires) {
			delete(r.slots, id)
			http.Error(w, "no such slot", http.StatusNotFound)
			return
		}
		if s.data != nil {
			http.Error(w, "already filled", http.StatusConflict)
			return
		}
		s.data = body
		w.WriteHeader(http.StatusNoContent)

	case http.MethodGet: // receiver collects; deleted on delivery
		r.mu.Lock()
		defer r.mu.Unlock()
		s, ok := r.slots[id]
		if !ok || time.Now().After(s.expires) {
			delete(r.slots, id)
			http.Error(w, "no such slot", http.StatusNotFound)
			return
		}
		if s.data == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		data := s.data
		delete(r.slots, id)
		w.Header().Set("Content-Type", "text/plain")
		w.Write(data)

	case http.MethodDelete: // receiver cancelled
		r.mu.Lock()
		delete(r.slots, id)
		r.mu.Unlock()
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, POST, PUT, DELETE")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func dropPrivileges() {
	if os.Geteuid() != 0 {
		return
	}
	// 65534 is "nobody" on Alpine and Debian
	if err := syscall.Setgid(65534); err != nil {
		log.Fatalf("setgid: %v", err)
	}
	if err := syscall.Setuid(65534); err != nil {
		log.Fatalf("setuid: %v", err)
	}
}

func main() {
	addr := os.Getenv("PAGERR_RELAY_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8081"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	dropPrivileges()

	r := newRelay()
	go func() {
		for range time.Tick(20 * time.Second) {
			r.sweep()
		}
	}()

	srv := &http.Server{
		Handler:           http.HandlerFunc(r.handle),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    8 << 10,
	}
	log.Printf("pagerr-relay listening on %s", addr)
	log.Fatal(srv.Serve(ln))
}
