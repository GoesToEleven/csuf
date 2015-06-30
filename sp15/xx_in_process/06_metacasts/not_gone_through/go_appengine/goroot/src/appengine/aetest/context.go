// Copyright 2013 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

/*
Package aetest provides an appengine.Context for use in tests.

An example test file:

	package foo_test

	import (
		"testing"

		"appengine/memcache"
		"appengine/aetest"
	)

	func TestFoo(t *testing.T) {
		c, err := aetest.NewContext(nil)
		if err != nil {
			t.Fatal(err)
		}
		defer c.Close()

		it := &memcache.Item{
			Key:   "some-key",
			Value: []byte("some-value"),
		}
		err = memcache.Set(c, it)
		if err != nil {
			t.Fatalf("Set err: %v", err)
		}
		it, err = memcache.Get(c, "some-key")
		if err != nil {
			t.Fatalf("Get err: %v; want no error", err)
		}
		if g, w := string(it.Value), "some-value" ; g != w {
			t.Errorf("retrieved Item.Value = %q, want %q", g, w)
		}
	}

The environment variable APPENGINE_DEV_APPSERVER specifies the location of the
dev_appserver.py executable to use. If unset, the system PATH is consulted.
*/
package aetest

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"appengine"
	"appengine/user"
	"appengine_internal"
	"github.com/golang/protobuf/proto"

	basepb "appengine_internal/base"
	remoteapipb "appengine_internal/remote_api"
)

// Context is an appengine.Context that sends all App Engine API calls to an
// instance of the API server.
type Context interface {
	appengine.Context

	// Login causes the context to act as the given user.
	Login(*user.User)
	// Logout causes the context to act as a logged-out user.
	Logout()
	// Close kills the child api_server.py process,
	// releasing its resources.
	io.Closer
}

// NewContext launches an instance of api_server.py and returns a Context
// that delegates all App Engine API calls to that instance.
// If opts is nil the default values are used.
func NewContext(opts *Options) (Context, error) {
	inst, err := NewInstance(opts)
	if err != nil {
		return nil, err
	}

	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		return nil, err
	}

	c := appengine.NewContext(req).(*context)
	return &singleContext{c}, nil
}

// context implements appengine.Context by delegating all Context calls to
// the provided instance.
type context struct {
	req      *http.Request
	instance Instance
	session  string
}

func (c *context) AppID() string               { return c.instance.appID() }
func (c *context) Request() interface{}        { return c.req }
func (c *context) FullyQualifiedAppID() string { return "dev~" + c.instance.appID() }

func (c *context) logf(level, format string, args ...interface{}) {
	log.Printf(level+": "+format, args...)
}

func (c *context) Debugf(format string, args ...interface{})    { c.logf("DEBUG", format, args...) }
func (c *context) Infof(format string, args ...interface{})     { c.logf("INFO", format, args...) }
func (c *context) Warningf(format string, args ...interface{})  { c.logf("WARNING", format, args...) }
func (c *context) Errorf(format string, args ...interface{})    { c.logf("ERROR", format, args...) }
func (c *context) Criticalf(format string, args ...interface{}) { c.logf("CRITICAL", format, args...) }

var errTimeout = &appengine_internal.CallError{
	Detail:  "Deadline exceeded",
	Code:    11, // CANCELED
	Timeout: true,
}

// postWithTimeout issues a POST to the specified URL with a given timeout.
func postWithTimeout(url, bodyType string, body io.Reader, timeout time.Duration) (b []byte, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	tr := &http.Transport{}
	client := &http.Client{
		Transport: tr,
	}
	if timeout != 0 {
		var canceled int32 // atomic; set to 1 if canceled
		t := time.AfterFunc(timeout, func() {
			atomic.StoreInt32(&canceled, 1)
			tr.CancelRequest(req)
		})
		defer t.Stop()
		defer func() {
			// Check to see whether the call was canceled.
			if atomic.LoadInt32(&canceled) != 0 {
				err = errTimeout
			}
		}()
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func call(service, method string, data []byte, apiAddress, requestID string, timeout time.Duration) ([]byte, error) {
	req := &remoteapipb.Request{
		ServiceName: proto.String(service),
		Method:      proto.String(method),
		Request:     data,
		RequestId:   proto.String(requestID),
	}

	buf, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	body, err := postWithTimeout(apiAddress, "application/octet-stream", bytes.NewReader(buf), timeout)
	if err != nil {
		return nil, err
	}

	res := &remoteapipb.Response{}
	err = proto.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	if ae := res.ApplicationError; ae != nil {
		// All Remote API application errors are API-level failures.
		return nil, &appengine_internal.APIError{Service: service, Detail: *ae.Detail, Code: *ae.Code}
	}
	return res.Response, nil
}

// Call is an implementation of appengine.Context's Call that delegates
// to a child api_server.py instance.
func (c *context) Call(service, method string, in, out appengine_internal.ProtoMessage, opts *appengine_internal.CallOptions) error {
	if service == "__go__" && (method == "GetNamespace" || method == "GetDefaultNamespace") {
		out.(*basepb.StringProto).Value = proto.String("")
		return nil
	}
	data, err := proto.Marshal(in)
	if err != nil {
		return err
	}
	var d time.Duration
	if opts != nil && opts.Timeout != 0 {
		d = opts.Timeout
	}
	res, err := call(service, method, data, c.instance.url(), c.session, d)
	if err != nil {
		return err
	}
	return proto.Unmarshal(res, out)
}

// singleContext is an implementation of context for an instance which
// only has a single context (namely, this one).
type singleContext struct {
	*context
}

func (c *singleContext) Login(u *user.User) {
	Login(u, c.context.req)
}

func (c *singleContext) Logout() {
	Logout(c.context.req)
}

func (c *singleContext) Close() error {
	return c.context.instance.Close()
}
