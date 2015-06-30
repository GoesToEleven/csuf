package main

import (
    "log"
    "net/http"
    "time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(time.RFC1123)
    w.Write([]byte("The time is: " + tm))
}

func main() {
    mux := http.NewServeMux()

    // Convert the timeHandler function to a HandleFunc type
    // http://golang.org/pkg/net/http/#HandlerFunc
    th := http.HandlerFunc(timeHandler)
    // And add it to the ServeMux
    mux.Handle("/time", th)

    log.Println("Listening beluga ...")
    http.ListenAndServe(":3000", mux)
}


/*
Processing HTTP requests with Go is primarily about two things:
(1) ServeMux aka Request Router aka MultiPlexor
(2) Handlers

Any function which has the signature func(http.ResponseWriter, *http.Request)
can be converted into a HandlerFunc type.
*/