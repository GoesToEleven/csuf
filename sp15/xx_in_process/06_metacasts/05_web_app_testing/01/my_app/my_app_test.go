package my_app_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/goestoeleven/mc/05_web_app_testing/01/my_app"
)

func TestHomePageHandler(t *testing.T) {
	mux := my_app.NewMux()
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(res, req)

	if res.Code != 200 {
		t.Errorf("!= 200")
	}

	if res.Body.String() != "Hello, World!" {
		t.Errorf("Hello, World!")
	}
}

func TestWelcomeByNameHandler_WithoutName(t *testing.T) {
	mux := my_app.NewMux()
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	mux.ServeHTTP(res, req)

	if res.Code != 200 {
		t.Error("!= 200")
	}

	if res.Body.String() != "Hello, World!" {
		t.Error("Hello, World!")
	}
}

func TestWelcomeByNameHandler_WithName(t *testing.T) {
	mux := my_app.NewMux()
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello?name=Todd", nil)
	mux.ServeHTTP(res, req)

	if res.Code != 200 {
		t.Error("!= 200")
	}

	if res.Body.String() != "Hello, Todd!" {
		t.Error("Hello, World!")
	}
}

func TestJsonHandler(t *testing.T) {
	mux := my_app.NewMux()
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/json", strings.NewReader(`{"first_name":"Todd","last_name":"McLeod"}`))
	mux.ServeHTTP(res, req)

	if res.Code != 201 {
		t.Error("!= 201")
	}

	if res.Header().Get("Content-Type") != "application/json" {
		t.Error("application/json")
	}

	user := new(my_app.User)
	json.NewDecoder(res.Body).Decode(user)

	if user.Id != 1 {
		t.Error("!= 1")
	}

	if user.FirstName != "Tod" {
		t.Error("!= Todd")
	}

	if user.LastName != "McLeod" {
		t.Errorf("%s != McLeod", user.LastName)
	}
}
