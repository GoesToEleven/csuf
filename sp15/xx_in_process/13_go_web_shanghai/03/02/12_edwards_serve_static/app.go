package main

import (
    "log"
    "net/http"
)

func main() {
    // "static" folder is our root
    fs := http.FileServer(http.Dir("static"))

    // http://localhost:3000/example.html
    // serves "static/example.html"
    http.Handle("/", fs)

    // http://localhost:3000/static/example.html
    // serves "static/example.html"
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // what URL serves: (answer below)
    // serves "static/deeper/deep.html"
    http.Handle("/deeper/", fs)


    log.Println("Listening orangutan ...")
    http.ListenAndServe(":3000", nil)
}

// this doesn't start from within webstorm
// this DOES START from the COMMAND LINE

/*
answer below ...





















http://localhost:3000/deeper/deep.html
*/