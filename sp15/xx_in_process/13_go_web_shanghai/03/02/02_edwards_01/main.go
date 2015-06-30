package main

import (
    "net/http"
    "log"
)

func main() {
    mux := http.NewServeMux()

    rh := http.RedirectHandler("http://www.google.com", 307)
    mux.Handle("/foo", rh)

    log.Println("Listening dog ...")
    http.ListenAndServe(":3000", mux)
}


/*
from:
http://www.alexedwards.net/blog/a-recap-of-request-handling

Processing HTTP requests with Go is primarily about two things:
(1) ServeMux aka Request Router aka MultiPlexor
(2) Handlers

A SERVEMUX is essentially a HTTP request router (or multiplexor).
---- ServeMux = Request Router = Multiplexor ----
It compares incoming requests against a list of predefined URL paths,
and calls the associated handler for the path whenever a match is found.

HANDLERS are responsible for writing response headers and bodies.
Almost any object can be a handler, so long as it satisfies the http.Handler interface.
In lay terms, that simply means it must have a ServeHTTP method with the following signature:

ServeHTTP(http.ResponseWriter, *http.Request)


*/