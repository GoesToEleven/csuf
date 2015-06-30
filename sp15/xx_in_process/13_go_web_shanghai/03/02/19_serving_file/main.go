package main

import (
    "net/http"
    "path"
)

func main() {
    http.HandleFunc("/", foo)
    http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
    // Assuming you want to serve a photo at 'images/foo.png'
    fp := path.Join("images", "foo.png")
    http.ServeFile(w, r, fp)
}

/*
$ curl -I localhost:3000

HTTP/1.1 200 OK
Accept-Ranges: bytes
Content-Length: 236717
Content-Type: image/png
Last-Modified: Thu, 10 Oct 2013 22:23:26 GMT

*/

// http://www.alexedwards.net/blog/golang-response-snippets