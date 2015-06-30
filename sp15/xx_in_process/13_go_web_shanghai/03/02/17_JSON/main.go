package main

import (
    "encoding/json"
    "net/http"
)

type Profile struct {
    Name    string
    Hobbies []string
}

func main() {
    http.HandleFunc("/", foo)
    http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
    profile := Profile{"Alex", []string{"snowboarding", "programming"}}

    js, err := json.Marshal(profile)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

/*
$ curl -i localhost:3000

HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 56

{"Name":"Alex",Hobbies":["snowboarding","programming"]}

*/

// http://www.alexedwards.net/blog/golang-response-snippets
