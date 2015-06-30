// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package logviewer

import (
	"encoding/base64"
	"errors"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"appengine"
	"appengine/log"
)

// appLevel holds data about a given application logging level.
type appLevel struct {
	Value  int
	Name   string
	Logger func(appengine.Context, string, ...interface{})
}

var (
	// funcMap holds a map of functions needed in the template environment.
	funcMap = map[string]interface{}{
		"levelInitial": levelInitial,
	}
	appLevels = []appLevel{
		{0, "Debug", appengine.Context.Debugf},
		{1, "Info", appengine.Context.Infof},
		{2, "Warning", appengine.Context.Warningf},
		{3, "Error", appengine.Context.Errorf},
		{4, "Critical", appengine.Context.Criticalf},
	}
	mainPageTmpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("log.html"))
)

func init() {
	http.HandleFunc("/", query)
	http.HandleFunc("/post", post)
}

// levelInitial converts an integer application log level to a one-letter
// string representing the level (e.g., "D" for Debug, "E" for Error).
func levelInitial(l int) (string, error) {
	if l < len(appLevels) {
		return appLevels[l].Name[:1], nil
	}
	return "", errors.New("Out of range application log level")
}

// query implements a basic HTML form to examine logs and format the output
// using Go templates.
func query(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Create a log.Query object; by default we request application logs,
	// but leave everything else at its zero value.
	query := &log.Query{AppLogs: true}

	// Parse the URL parameters, handling absent or invalid values.
	version := r.FormValue("version")
	if len(version) != 0 {
		query.Versions = []string{version}
	} else {
		// Compute the current major version ID for the form.
		version = appengine.VersionID(c)
		if i := strings.Index(version, "."); i > -1 {
			version = version[:i]
		}
	}
	level, err := strconv.Atoi(r.FormValue("level"))
	if err == nil && level > -1 && level < len(appLevels) {
		query.ApplyMinLevel = true
		query.MinLevel = level
	} else {
		level = -1
	}
	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil {
		count = 25
	}
	offset, err := base64.URLEncoding.DecodeString(r.FormValue("offset"))
	if err == nil && len(offset) > 0 {
		query.Offset = offset
	}

	// Set up a map of untyped values to pass to the log template.  Our
	// interpretations of the URL parameters are here to populate the form.
	data := struct {
		AppLevels []appLevel
		Version   string
		Level     int
		Count     int
		Error     error
		Next      string
		Logs      []*log.Record
	}{
		AppLevels: appLevels, // We'll build the form from appLevels as well.
		Version:   version,
		Level:     level,
		Count:     count,
	}

	// Run the query and read at most count records.
	for results := query.Run(c); len(data.Logs) < count; {
		record, err := results.Next()
		if err != nil {
			if err != log.Done {
				c.Errorf("Failed to retrieve log: %v", err)
				data.Error = err
			}
			break
		}
		data.Logs = append(data.Logs, record)
	}
	if len(data.Logs) == count {
		v := r.URL.Query()
		offset := data.Logs[len(data.Logs)-1].Offset
		v.Set("offset", base64.URLEncoding.EncodeToString(offset))
		data.Next = v.Encode()
	}

	// Execute the main page template and return the response.
	if err := mainPageTmpl.ExecuteTemplate(w, "log.html", data); err != nil {
		c.Errorf("Failed to execute template: %v", err)
	}
}

// post provides an easy way to add application logs to this version's stream.
func post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	level, err := strconv.Atoi(r.FormValue("level"))
	message := r.FormValue("message")
	if err == nil && level < len(appLevels) { // On error just skip logging.
		appLevels[level].Logger(c, "%s", message)
	}

	// Our redirect utilizes some Javascript on the page to pass the URL
	// parameters to us so we can repopulate the user's existing search.
	http.Redirect(w, r, "/"+r.FormValue("search"), http.StatusFound)
}
