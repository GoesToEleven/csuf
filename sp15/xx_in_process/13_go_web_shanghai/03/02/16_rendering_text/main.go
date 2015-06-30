package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", foo)
    http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
}

/*
curl -i localhost:3000

HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Content-Length: 2

OK
*/

// http://www.alexedwards.net/blog/golang-response-snippets