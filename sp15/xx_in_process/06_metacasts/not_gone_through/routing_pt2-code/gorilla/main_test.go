package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Routing(t *testing.T) {
	a := assert.New(t)

	ts := httptest.NewServer(NewMux())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users")
	a.NoError(err)
	body, err := ioutil.ReadAll(res.Body)
	a.Equal(string(body), "Users Index")

	res, err = http.Get(ts.URL + "/users/42")
	a.NoError(err)
	body, err = ioutil.ReadAll(res.Body)
	a.Equal(string(body), "Users Show: 42")

	res, err = http.Post(ts.URL+"/users", "plain/text", strings.NewReader("Hello!"))
	a.NoError(err)
	body, err = ioutil.ReadAll(res.Body)
	a.Equal(string(body), "Users Create: Hello!")

	res, err = http.Get(ts.URL + "/posts")
	a.NoError(err)
	body, err = ioutil.ReadAll(res.Body)
	a.Equal(string(body), "POSTS!")
}
