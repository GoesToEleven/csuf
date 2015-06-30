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

    lp := path.Join("templates", "layout.html")
    fp := path.Join("templates", "index.html")

    // Note that the layout file must be the first parameter in ParseFiles
    tmpl, err := template.ParseFiles(lp, fp)
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
Date: Sat, 25 Apr 2015 07:02:51 GMT
Content-Length: 168
Content-Type: text/html; charset=utf-8

<html>
<head>
    <title>An example layout</title>
</head>
<body>

<h1>Hello Alex</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.</p>

</body>
</html>
*/