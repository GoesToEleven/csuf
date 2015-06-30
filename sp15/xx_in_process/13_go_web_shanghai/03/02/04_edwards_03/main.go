package main

import (
    "net/http"
    "time"
    "log"
)

type timeHandler struct {
    format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(th.format)
    w.Write([]byte("The time is: " + tm))
}

func main() {
    mux := http.NewServeMux()

    th := &timeHandler{format: time.RFC1123}
    mux.Handle("/time", th)

    log.Println("Listening, yo ...")
    http.ListenAndServe(":3000", mux)
}

/*
Processing HTTP requests with Go is primarily about two things:
(1) ServeMux aka Request Router aka MultiPlexor
(2) Handlers
*/
