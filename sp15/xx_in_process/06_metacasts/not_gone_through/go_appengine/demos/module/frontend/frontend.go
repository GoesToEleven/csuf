// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package frontend implements a frontend to the example memcache module.
// TODO: Demonstrate sharding across multiple instances.
package frontend

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"

	"appengine"
	"appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", handleFront)
}

func handleFront(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	c := appengine.NewContext(r)

	addr, err := memcacheAddr(c)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed getting module address: %v", err), 500)
		return
	}

	key, value := r.FormValue("key"), r.FormValue("value")
	m := map[string]interface{}{
		"Key":             key,
		"Value":           value,
		"MemcacheAddress": addr,
	}

	switch r.Method {
	case "GET":
		m["GetValue"], m["Error"] = get(c, key)
	case "POST":
		m["SetMessage"], m["Error"] = set(c, key, value)
	}

	if err := page.Execute(w, m); err != nil {
		c.Errorf("Template execution failed: %v", err)
	}
}

func memcacheAddr(c appengine.Context) (string, error) {
	// Use the load-balanced hostname for the "memcache" module.
	hostname, err := appengine.ModuleHostname(c, "memcache", "", "")
	if err != nil {
		return "", err
	}
	return "http://" + hostname, nil
}

func get(c appengine.Context, key string) ([]byte, error) {
	addr, err := memcacheAddr(c)
	if err != nil {
		return nil, err
	}
	u := addr + "/memcache/get?key=" + url.QueryEscape(key)
	resp, err := urlfetch.Client(c).Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func set(c appengine.Context, key, value string) (string, error) {
	addr, err := memcacheAddr(c)
	if err != nil {
		return "", err
	}
	u := addr + "/memcache/set"
	resp, err := urlfetch.Client(c).PostForm(u, url.Values{
		"key":   {key},
		"value": {value},
	})
	return resp.Status, err
}

var page = template.Must(template.New("front").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>Memcache frontend</title>
	<style type="text/css">
	.box {
		border: 1px solid black;
		margin: 0.5em;
		padding: 1em;
		width: 20em;
	}
	.error {
		color: red;
		font-weight: bold;
	}
	</style>
</head>
<body>
<h1>Memcache frontend</h1>
	{{with .Error}}
	<span class="error">{{.}}</span>
	{{end}}

	<div class="box">
	<h3>Get</h3>
	<form method=GET action="/">
		<label for="key">Key:</label>
		<input type="text" name="key" id="key" value="{{.Key}}" /><br />
		<input type="submit" /><br />
		{{with .GetValue}}
		<b>{{printf "%s" .}}</b>
		{{end}}
	</form>
	</div>

	<div class="box">
	<h3>Set</h3>
	<form method=POST action="/">
		<label for="key">Key:</label>
		<input type="text" name="key" id="key" value="{{.Key}}" /><br />
		<label for="value">Value:</label>
		<input type="text" name="value" id="value" value="{{.Value}}" /><br />
		<input type="submit" /><br />
		{{with .SetMessage}}
		<b>{{.}}</b>
		{{end}}
	</form>
	</div>

	<p>Memcache address: <code>{{.MemcacheAddress}}</code></p>
</body>
</html>
`))
