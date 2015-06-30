// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package guestbook

import (
	"html/template"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

func init() {
	http.HandleFunc("/", handleMainPage)
	http.HandleFunc("/sign", handleSign)
}

// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c appengine.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}

var mainPage = template.Must(template.New("guestbook").Parse(
	`<html><body>
{{range .}}
{{with .Author}}<b>{{.}}</b>{{else}}An anonymous person{{end}}
on <em>{{.Date.Format "3:04pm, Mon 2 Jan"}}</em>
wrote <blockquote>{{.Content}}</blockquote>
{{end}}
<form action="/sign" method="post">
<div><textarea name="content" rows="3" cols="60"></textarea></div>
<div><input type="submit" value="Sign Guestbook"></div>
</form></body></html>
`))

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "GET requests only", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	c := appengine.NewContext(r)
	q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(10)
	var gg []*Greeting
	if _, err := q.GetAll(c, &gg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := mainPage.Execute(w, gg); err != nil {
		c.Errorf("%v", err)
	}
}

func handleSign(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "POST requests only", http.StatusMethodNotAllowed)
		return
	}
	c := appengine.NewContext(r)
	g := &Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}
	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}
	key := datastore.NewIncompleteKey(c, "Greeting", guestbookKey(c))
	if _, err := datastore.Put(c, key, g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
