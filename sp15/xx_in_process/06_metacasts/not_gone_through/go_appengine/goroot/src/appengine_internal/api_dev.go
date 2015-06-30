// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package appengine_internal

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	basepb "appengine_internal/base"
	lpb "appengine_internal/log"
	"appengine_internal/remote_api"
	rpb "appengine_internal/runtime_config"
	"github.com/golang/protobuf/proto"
)

// IsDevAppServer returns whether the App Engine app is running in the
// development App Server.
func IsDevAppServer() bool {
	return true
}

// serveHTTP serves App Engine HTTP requests.
func serveHTTP() {
	// The development server reads the HTTP port that the server is listening to
	// from stdout. We listen on 127.0.0.1:0 to avoid firewall restrictions.
	conn, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal("appengine: couldn't listen to TCP socket: ", err)
	}
	port := conn.Addr().(*net.TCPAddr).Port

	fmt.Fprintln(os.Stdout, port)
	os.Stdout.Close()

	err = http.Serve(conn, http.HandlerFunc(handleFilteredHTTP))
	if err != nil {
		log.Fatal("appengine: ", err)
	}
}

func init() {
	// If the user's application has a transitive dependency on appengine_internal
	// then this init will be called before any user code.
	// Read configuration from stdin when the application is being run by
	// devappserver2. The user application should not be reading from stdin.
	if os.Getenv("RUN_WITH_DEVAPPSERVER") != "1" {
		log.Print("appengine: not running under devappserver2; using some default configuration")
		return
	}
	c := readConfig(os.Stdin)
	instanceConfig.AppID = string(c.AppId)
	instanceConfig.APIHost = c.GetApiHost()
	instanceConfig.APIPort = int(*c.ApiPort)
	instanceConfig.VersionID = string(c.VersionId)
	instanceConfig.InstanceID = *c.InstanceId
	instanceConfig.Datacenter = *c.Datacenter

	apiAddress = fmt.Sprintf("http://%s:%d", instanceConfig.APIHost, instanceConfig.APIPort)
}

func handleFilteredHTTP(w http.ResponseWriter, r *http.Request) {
	// Patch up RemoteAddr so it looks reasonable.
	if addr := r.Header.Get("X-Appengine-Remote-Addr"); addr != "" {
		r.RemoteAddr = addr
	} else {
		// Should not normally reach here, but pick
		// a sensible default anyway.
		r.RemoteAddr = "127.0.0.1"
	}

	// Create a private copy of the Request that includes headers that are
	// private to the runtime and strip those headers from the request that the
	// user application sees.
	creq := *r
	r.Header = make(http.Header)
	for name, values := range creq.Header {
		if !strings.HasPrefix(name, "X-Appengine-Dev-") {
			r.Header[name] = values
		}
	}
	ctxsMu.Lock()
	ctxs[r] = &httpContext{req: &creq}
	ctxsMu.Unlock()

	http.DefaultServeMux.ServeHTTP(w, r)

	ctxsMu.Lock()
	delete(ctxs, r)
	ctxsMu.Unlock()
}

var (
	apiAddress    string
	apiHTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	ctxsMu sync.Mutex
	ctxs   = make(map[*http.Request]context)

	instanceConfig = struct {
		AppID      string
		VersionID  string
		InstanceID string
		Datacenter string
		APIHost    string
		APIPort    int
	}{
		// Default configuration for when this file is loaded outside the context
		// of devappserver2.
		AppID:      "dev~my~app",
		VersionID:  "1.2345",
		InstanceID: "deadbeef",
		Datacenter: "us1",
		APIHost:    "localhost",
		APIPort:    1,
	}
)

func readConfig(r io.Reader) *rpb.Config {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal("appengine: could not read from stdin: ", err)
	}

	if len(raw) == 0 {
		log.Fatal("appengine: no config provided on stdin")
	}

	b := make([]byte, base64.StdEncoding.DecodedLen(len(raw)))
	n, err := base64.StdEncoding.Decode(b, raw)
	if err != nil {
		log.Fatal("appengine: could not base64 decode stdin: ", err)
	}
	config := &rpb.Config{}

	err = proto.Unmarshal(b[:n], config)
	if err != nil {
		log.Fatal("appengine: could not decode runtime_config: ", err)
	}
	return config
}

var errTimeout = &CallError{
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
	if timeout != 0 {
		if tr, ok := apiHTTPClient.Transport.(*http.Transport); ok {
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
	}
	resp, err := apiHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func call(service, method string, data []byte, requestID string, timeout time.Duration) ([]byte, error) {
	req := &remote_api.Request{
		ServiceName: &service,
		Method:      &method,
		Request:     data,
		RequestId:   &requestID,
	}

	buf, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	body, err := postWithTimeout(apiAddress, "application/octet-stream", bytes.NewReader(buf), timeout)
	if err != nil {
		return nil, err
	}

	res := &remote_api.Response{}
	err = proto.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	if ae := res.ApplicationError; ae != nil {
		// All Remote API application errors are API-level failures.
		return nil, &APIError{Service: service, Detail: *ae.Detail, Code: *ae.Code}
	}
	return res.Response, nil
}

// context echos the public appengine.Context interface.
type context interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Criticalf(format string, args ...interface{})
	Call(service, method string, in, out ProtoMessage, opts *CallOptions) error
	FullyQualifiedAppID() string
	Request() interface{}
}

// httpContext represents the context of an in-flight HTTP request.
// It implements the appengine.Context interface.
type httpContext struct {
	req *http.Request
}

func NewContext(req *http.Request) context {
	ctxsMu.Lock()
	defer ctxsMu.Unlock()
	c := ctxs[req]

	if c == nil {
		// Someone passed in an http.Request that is not in-flight.
		// We panic here rather than panicking at a later point
		// so that backtraces will be more sensible.
		log.Panic("appengine: NewContext passed an unknown http.Request")
	}
	return c
}

// RegisterTestContext associates a test context with the given HTTP request,
// returning a closure to delete the association. It should only be used by the
// aetest package, and never directly. It is only available in the SDK.
func RegisterTestContext(req *http.Request, c context) func() {
	ctxsMu.Lock()
	defer ctxsMu.Unlock()
	if _, ok := ctxs[req]; ok {
		log.Panic("req already associated with context")
	}
	ctxs[req] = c

	return func() {
		ctxsMu.Lock()
		delete(ctxs, req)
		ctxsMu.Unlock()
	}
}

func (c *httpContext) Call(service, method string, in, out ProtoMessage, opts *CallOptions) error {
	if service == "__go__" {
		if method == "GetNamespace" {
			out.(*basepb.StringProto).Value = proto.String(c.req.Header.Get("X-AppEngine-Current-Namespace"))
			return nil
		}
		if method == "GetDefaultNamespace" {
			out.(*basepb.StringProto).Value = proto.String(c.req.Header.Get("X-AppEngine-Default-Namespace"))
			return nil
		}
	}
	if f, ok := apiOverrides[struct{ service, method string }{service, method}]; ok {
		return f(in, out, opts)
	}
	data, err := proto.Marshal(in)
	if err != nil {
		return err
	}

	requestID := c.req.Header.Get("X-Appengine-Dev-Request-Id")
	var d time.Duration
	if opts != nil && opts.Timeout != 0 {
		d = opts.Timeout
	}
	res, err := call(service, method, data, requestID, d)
	if err != nil {
		return err
	}
	return proto.Unmarshal(res, out)
}

func (c *httpContext) Request() interface{} {
	return c.req
}

func (c *httpContext) logf(level int64, levelName, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	s = strings.TrimRight(s, "\n") // Remove any trailing newline characters.
	log.Println(levelName + ": " + s)

	// Truncate long log lines.
	const maxLogLine = 8192
	if len(s) > maxLogLine {
		suffix := fmt.Sprintf("...(length %d)", len(s))
		s = s[:maxLogLine-len(suffix)] + suffix
	}

	buf, err := proto.Marshal(&lpb.UserAppLogGroup{
		LogLine: []*lpb.UserAppLogLine{
			{
				TimestampUsec: proto.Int64(time.Now().UnixNano() / 1e3),
				Level:         proto.Int64(level),
				Message:       proto.String(s),
			}}})
	if err != nil {
		log.Printf("appengine_internal.flushLog: failed marshaling AppLogGroup: %v", err)
		return
	}

	req := &lpb.FlushRequest{
		Logs: buf,
	}
	res := &basepb.VoidProto{}
	if err := c.Call("logservice", "Flush", req, res, nil); err != nil {
		log.Printf("appengine_internal.flushLog: failed Flush RPC: %v", err)
	}
}

func (c *httpContext) Debugf(format string, args ...interface{}) { c.logf(0, "DEBUG", format, args...) }
func (c *httpContext) Infof(format string, args ...interface{})  { c.logf(1, "INFO", format, args...) }
func (c *httpContext) Warningf(format string, args ...interface{}) {
	c.logf(2, "WARNING", format, args...)
}
func (c *httpContext) Errorf(format string, args ...interface{}) { c.logf(3, "ERROR", format, args...) }
func (c *httpContext) Criticalf(format string, args ...interface{}) {
	c.logf(4, "CRITICAL", format, args...)
}

// FullyQualifiedAppID returns the fully-qualified application ID.
// This may contain a partition prefix (e.g. "s~" for High Replication apps),
// or a domain prefix (e.g. "example.com:").
func (c *httpContext) FullyQualifiedAppID() string {
	return instanceConfig.AppID
}
