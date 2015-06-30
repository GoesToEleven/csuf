package main

import (
    "fmt"
    "net/http"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there.")
}

func main() {

    http.HandleFunc("/", handler)

    log.Println("Listening cowboy ...")
    http.ListenAndServe(":8080", nil)
}

/*
Processing HTTP requests with Go is primarily about two things:
(1) ServeMux aka Request Router aka MultiPlexor
(2) Handlers

HandleFunc
converting a function to a HandlerFunc type and then adding it to a ServeMux is so common
Go provides a shortcut: the ServeMux.HandleFunc method
*/
