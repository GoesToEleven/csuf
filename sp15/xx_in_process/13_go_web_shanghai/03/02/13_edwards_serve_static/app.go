package main

import (
"html/template"
"log"
"net/http"
"path"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", serveTemplate)

    log.Println("Listening tractor ...")
    http.ListenAndServe(":3000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
    lp := path.Join("templates", "layout.html")
    fp := path.Join("templates", r.URL.Path)

    tmpl, _ := template.ParseFiles(lp, fp)
    tmpl.ExecuteTemplate(w, "layout", nil)
}

// this doesn't start from within webstorm
// this DOES START from the COMMAND LINE

/*
In the serveTemplate function, we build paths to the layout file and the template file
corresponding with the request. Rather than manual concatenation we use Join, which has
the advantage of cleaning the path to help prevent directory traversal attacks.

to run:
http://localhost:3000/example.html
http://localhost:3000/static/example.html
*/