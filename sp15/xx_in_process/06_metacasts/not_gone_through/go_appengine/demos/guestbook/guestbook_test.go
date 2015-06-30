// Copyright 2014 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package guestbook

import (
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"appengine"
	"appengine/aetest"
	"appengine/datastore"
	"appengine/user"
)

func TestSign(t *testing.T) {
	now := time.Now()
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	testCases := []struct {
		method   string
		content  string
		user     *user.User
		code     int
		greeting *Greeting
	}{
		{
			method: "GET",
			code:   405,
		},
		{
			method:   "POST",
			content:  "Normal post",
			code:     303,
			greeting: &Greeting{Content: "Normal post"},
		},
		{
			method:   "POST",
			content:  "Post with user",
			user:     &user.User{Email: "josmith@gmail.com"},
			code:     303,
			greeting: &Greeting{Content: "Post with user", Author: "josmith@gmail.com"},
		},
	}

	for _, tt := range testCases {
		data := url.Values{
			"content": []string{tt.content},
		}
		req, err := inst.NewRequest(tt.method, "/post", strings.NewReader(data.Encode()))
		if err != nil {
			t.Errorf("inst.NewRequest failed: %v", err)
			continue
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		if tt.user != nil {
			aetest.Login(tt.user, req)
		}

		resp := httptest.NewRecorder()
		handleSign(resp, req)

		if resp.Code != tt.code {
			t.Errorf("Got response code %d; want %d; body:\n%s", resp.Code, tt.code, resp.Body.String())
		}

		// Check the lastest greeting against our expectation.
		c := appengine.NewContext(req)
		q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(1)
		var g Greeting
		_, err = q.Run(c).Next(&g)
		if err == datastore.Done {
			if tt.greeting != nil {
				t.Errorf("No greeting stored. Expected %v", tt.greeting)
			}
			continue
		}
		if err != nil {
			t.Errorf("Failed to fetch greeting: %v", err)
			continue
		}
		if tt.greeting == nil {
			if !g.Date.Before(now) {
				t.Errorf("Expected no new greeting, found: %v", g)
			}
			continue
		}
		if g.Date.Before(now) {
			t.Errorf("Greeting stored at %v, want at least %v", g.Date, now)
		}
		g.Date = time.Time{} // Zero out for comparisons.
		if !reflect.DeepEqual(g, *tt.greeting) {
			t.Errorf("Greetings don't match.\nGot:  %v\nWant: %v", g, *tt.greeting)
		}
	}
}
