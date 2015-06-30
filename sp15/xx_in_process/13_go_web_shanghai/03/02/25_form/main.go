package main

import (
    "github.com/julienschmidt/httprouter"
    "html/template"
    "log"
    "net/http"
    "github.com/goestoeleven/csuf/sp15/xx_in_process/06_metacasts/not_gone_through/go_appengine/goroot/src/fmt"
)

func main() {
    router := httprouter.New()
    router.GET("/", index)
    router.POST("/", send)
    router.GET("/confirmation", confirmation)

    log.Println("Listening bulldog ...")
    http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    render(w, "templates/index.html", nil)
}

func send(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//    email := params.ByName("email")
//    content := params.ByName("content")
//    fmt.Fprintln(w, email + " " + content)
}

func confirmation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    render(w, "templates/confirmation.html", nil)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
    tmpl, err := template.ParseFiles(filename)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}