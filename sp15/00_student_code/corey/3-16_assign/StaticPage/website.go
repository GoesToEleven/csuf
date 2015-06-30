package main

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(rw http.ResponseWriter, req *http.Request) {

}
