package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"appengine/aetest"
)

func TestAPI(t *testing.T) {
	assert := assert.New(t)

	c, err := aetest.NewContext(nil)
	assert.Nil(err)
	defer c.Close()

	inst, err := aetest.NewInstance(nil)
	assert.Nil(err)
	defer inst.Close()

	cookie := ""
	makeRequest := func(method string, endpoint string, data interface{}) *httptest.ResponseRecorder {
		var body io.Reader
		var contentType string
		if r, ok := data.(io.Reader); ok {
			body = r
		} else if data != nil {
			bs, _ := json.Marshal(data)
			body = bytes.NewReader(bs)
			contentType = "application/json"
		}
		req, err := inst.NewRequest(method, endpoint, body)
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Cookie", cookie)
		assert.Nil(err)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		sc := res.Header().Get("Set-Cookie")
		if sc != "" {
			sc = sc[:strings.Index(sc, ";")]
			cookie = sc
		}
		return res
	}

	// create user
	res := makeRequest("POST", "/api/users", map[string]string{
		"email":    "test@example.com",
		"password": "test",
	})
	assert.Equal(200, res.Code)

	// login as user
	res = makeRequest("POST", "/api/users/login", map[string]string{
		"email":    "test@example.com",
		"password": "test",
	})
	assert.Equal(200, res.Code)

	// upload file
	res = makeRequest("POST", "/api/files", strings.NewReader("Hello World"))
	assert.Equal(200, res.Code)
	var fileid string
	json.Unmarshal(res.Body.Bytes(), &fileid)

	// download file
	res = makeRequest("GET", "/api/files/"+fileid, nil)
	assert.Equal(200, res.Code)
	assert.Equal("Hello World", res.Body.String())

	// create document
	res = makeRequest("POST", "/api/documents", Document{
		Link:     "/path/to/file",
		Contents: "Hello World",
	})
	assert.Equal(200, res.Code)
	var doc Document
	json.Unmarshal(res.Body.Bytes(), &doc)
	assert.Equal("Hello World", doc.Contents)

	// get document
	res = makeRequest("GET", "/api/documents/"+doc.ID, nil)
	assert.Equal(200, res.Code)
	json.Unmarshal(res.Body.Bytes(), &doc)
	assert.Equal("Hello World", doc.Contents)

	// update document
	doc.Contents = "Hello World 2"
	res = makeRequest("PUT", "/api/documents/"+doc.ID, doc)
	assert.Equal(200, res.Code)
	json.Unmarshal(res.Body.Bytes(), &doc)
	assert.Equal("Hello World 2", doc.Contents)

	// list documents
	res = makeRequest("GET", "/api/documents", nil)
	assert.Equal(200, res.Code)
	var docs []Document
	json.Unmarshal(res.Body.Bytes(), &docs)
	assert.Equal(1, len(docs))
	assert.Equal("Hello World 2", docs[0].Contents)
}
