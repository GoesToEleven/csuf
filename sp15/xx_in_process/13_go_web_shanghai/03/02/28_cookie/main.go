package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/set", set)
    http.HandleFunc("/get", get)
    fmt.Println("Listening koala ...")
    http.ListenAndServe(":3000", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
    fm := []byte("This is a flashed message! Here is more content.")
    SetFlash(w, "message", fm)
}

func get(w http.ResponseWriter, r *http.Request) {
    fm, err := GetFlash(w, r, "message")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if fm == nil {
        fmt.Fprint(w, "No flash messages")
        return
    }
    fmt.Fprintf(w, "%s", fm)
}