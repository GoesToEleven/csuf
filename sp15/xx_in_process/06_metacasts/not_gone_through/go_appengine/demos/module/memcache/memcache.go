// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package memcache implements a simple memcache-alike.
package memcache

import (
	"io"
	"net/http"
	"strconv"
	"sync"

	"appengine"
)

func init() {
	http.HandleFunc("/_ah/start", handleStart)
	http.HandleFunc("/_ah/stop", handleStop)
}

var (
	mu    sync.Mutex
	cache map[string][]byte
)

func handleStart(w http.ResponseWriter, r *http.Request) {
	// This handler is executed when an instance is started.
	// If it responds with a HTTP 2xx or 404 response then it is ready to go.
	// Otherwise, the instance is terminated and restarted.
	// The instance will receive traffic after this handler returns.

	c := appengine.NewContext(r)
	http.HandleFunc("/memcache/get", handleGet)
	http.HandleFunc("/memcache/set", handleSet)
	cache = make(map[string][]byte)
	c.Infof("Memcache module started.")
	io.WriteString(w, "OK")
}

func handleStop(w http.ResponseWriter, r *http.Request) {
	// This handler is executed when an instance is being shut down.
	// It has 30s before it will be terminated.
	// When this is called, no new requests will reach the instance.

	c := appengine.NewContext(r)
	c.Infof("Memcache module stopped.")
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	key := r.FormValue("key")
	mu.Lock()
	value, ok := cache[key]
	mu.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNoContent)
		c.Infof("No data for key %q", key)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(value)))
	w.Write(value)
	c.Infof("Returned %q => %q", key, value)
}

func handleSet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	c := appengine.NewContext(r)
	key, value := r.FormValue("key"), r.FormValue("value")
	mu.Lock()
	cache[key] = []byte(value)
	mu.Unlock()

	c.Infof("Stored %q => %q", key, value)
}
