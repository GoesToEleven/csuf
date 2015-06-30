// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package later

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/delay"
)

func init() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// ignore /favicon.ico, etc.
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "I will run a function later.\n")

	c := appengine.NewContext(r)
	later.Call(c, "Please send my regards to your mother.")
}

// A function that can be executed later.
var later = delay.Func("later", func(c appengine.Context, msg string) {
	c.Infof("later, %q", msg)
})
