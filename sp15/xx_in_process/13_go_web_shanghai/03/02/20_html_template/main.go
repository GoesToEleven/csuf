package main

import (
    "html/template"
    "net/http"
    "path"
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

    fp := path.Join("templates", "index.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, profile); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


/*
$ curl -i localhost:3000

HTTP/1.1 200 OK
Content-Type: text/html; charset=utf-8
Content-Length: 84

<h1>Hello Alex</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.</p>

*/

// http://www.alexedwards.net/blog/golang-response-snippets