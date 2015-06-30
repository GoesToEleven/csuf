package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    "regexp"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()

        if len(r.Form["email"][0])==0 {
            fmt.Fprintf(w, "email was empty")
        }

        if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
            fmt.Println("no")
            fmt.Fprintf(w, "that is not a valid email")
        }else {
            fmt.Println("yes")
        }

        // logic part of log in
        fmt.Println("r.form:", r.Form)
        fmt.Println("email:", r.Form["email"])
        fmt.Fprintf(w, "email: %s", r.Form["email"][0])
    }
}

func main() {
    http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/login", login)
    log.Println("Listening swahili ...")
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}