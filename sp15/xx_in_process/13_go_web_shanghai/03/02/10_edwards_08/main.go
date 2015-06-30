package main

import (
    "log"
    "net/http"
    "time"
)

func timeHandler(format string) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        tm := time.Now().Format(format)
        w.Write([]byte("The time is: " + tm))
    }
    return http.HandlerFunc(fn)
}

func main() {
    // Note that we skip creating the ServeMux...

    var format string = time.RFC1123
    th := timeHandler(format)

    // We use http.Handle instead of mux.Handle...
    http.Handle("/time", th)

    log.Println("Listening...")
    // And pass nil as the handler to ListenAndServe.
    http.ListenAndServe(":3000", nil)
}

/*
Processing HTTP requests with Go is primarily about two things:
(1) ServeMux aka Request Router aka MultiPlexor
(2) Handlers

The DefaultServeMux

You've probably seen DefaultServeMux mentioned in lots of places,
from the simplest Hello World examples to the Go source code.

It took me a long time to realise it isn't anything special.
The DefaultServeMux is just a plain ol' ServeMux like we've already been using,
which gets instantiated by default when the HTTP package is used.

Here's the relevant line from the Go source:

var DefaultServeMux = NewServeMux()

The HTTP package provides a couple of shortcuts for working with the DefaultServeMux:

                    http.Handle

                    http.HandleFunc

These do exactly the same as their namesake functions we've already looked at,
with the difference that they add handlers to the DefaultServeMux
instead of one that you've created.

Additionally, ListenAndServe will fall back to using the DefaultServeMux
if no other handler is provided (that is, the second parameter is set to nil).

*/