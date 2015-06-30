package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", foo)
    http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", "A Go Web Server")
    w.WriteHeader(200)
}

/*
curl -i localhost:3000

HTTP/1.1 200 OK
Server: A Go Web Server
Date: Sat, 25 Apr 2015 06:26:20 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8

*/

// http://www.alexedwards.net/blog/golang-response-snippets