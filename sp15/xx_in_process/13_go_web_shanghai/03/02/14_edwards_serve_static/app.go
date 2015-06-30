package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
    "path"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.HandleFunc("/", serveTemplate)

    log.Println("Listening submarine ...")
    http.ListenAndServe(":3000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
    lp := path.Join("templates", "layout.html")
    fp := path.Join("templates", r.URL.Path)

    // Return a 404 if the template doesn't exist
    info, err := os.Stat(fp)
    if err != nil {
        if os.IsNotExist(err) {
            http.NotFound(w, r)
            return
        }
    }

    // Return a 404 if the request is for a directory
    if info.IsDir() {
        http.NotFound(w, r)
        return
    }

    tmpl, err := template.ParseFiles(lp, fp)
    if err != nil {
        // Log the detailed error
        log.Println(err.Error())
        // Return a generic "Internal Server Error" message
        http.Error(w, http.StatusText(500), 500)
        return
    }

    if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
        log.Println(err.Error())
        http.Error(w, http.StatusText(500), 500)
    }
}

// this doesn't start from within webstorm
// this DOES START from the COMMAND LINE

/*
Send a 404 response if the requested template doesn't exist.

Send a 404 response if the requested template path is a directory.

Send a 500 response if the template.ParseFiles or template.ExecuteTemplate
functions throw an error, and log the detailed error message.

to run:
http://localhost:3000/example.html
http://localhost:3000/static/example.html
*/