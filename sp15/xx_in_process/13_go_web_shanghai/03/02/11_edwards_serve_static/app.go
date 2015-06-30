package main

import (
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)

    log.Println("Listening platypus ...")
    http.ListenAndServe(":3000", nil)
}

// this doesn't start from within webstorm
// this DOES START from the COMMAND LINE