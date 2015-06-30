package main

import (
    "html/template"
    "net/http"
    "path"
    "bytes"
    "github.com/goestoeleven/csuf/sp15/xx_in_process/06_metacasts/not_gone_through/go_appengine/goroot/src/fmt"
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

    buf := new(bytes.Buffer)
    if err := tmpl.Execute(buf, profile); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    templateString := buf.String()
    fmt.Fprintln(w, templateString)
}


/*
$ curl -i localhost:3000

HTTP/1.1 200 OK
Date: Sat, 25 Apr 2015 06:59:08 GMT
Content-Length: 85
Content-Type: text/html; charset=utf-8

<h1>Hello Alex</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.</p>
*/

// http://www.alexedwards.net/blog/golang-response-snippets