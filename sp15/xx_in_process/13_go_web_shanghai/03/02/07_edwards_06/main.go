package main

import (
    "net/http"
    "log"
    "time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(time.RFC1123)
    w.Write([]byte("The time is: " + tm))
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/time", timeHandler)

    log.Println("Listening soliloquy ...")
    http.ListenAndServe(":3000", mux)
}

/*
Processing HTTP requests with Go is primarily about two things:
(1) ServeMux aka Request Router aka MultiPlexor
(2) Handlers

HandleFunc
converting a function to a HandlerFunc type and then adding it to a ServeMux is so common
Go provides a shortcut: the ServeMux.HandleFunc method
*/