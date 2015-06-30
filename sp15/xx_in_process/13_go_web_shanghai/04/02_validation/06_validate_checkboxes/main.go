package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
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

// Returns true if slices have a common element
func doSlicesIntersect(s1, s2 []string) bool {
    if s1 == nil || s2 == nil {
        return false
    }
    for _, str := range s1 {
        if isElementInSlice(str, s2) {
            return true
        }
    }
    return false
}

func isElementInSlice(str string, sl []string) bool {
    if sl == nil || str == "" {
        return false
    }
    for _, v := range sl {
        if v == str {
            return true
        }
    }
    return false
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()

        slice:=[]string{"football","basketball","tennis"}
        validInterest := doSlicesIntersect(slice, r.Form["interest"])

        if !validInterest {
            fmt.Fprintf(w, "that is not an interest")
        } else {
            fmt.Fprintf(w, "that is an interest")
        }

        // logic part of log in
        fmt.Println("r.form:", r.Form)
        fmt.Println("interest:", r.Form["interest"])
        fmt.Fprintf(w, "interest: %s", r.Form["interest"])
    }
}

func main() {
    http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/login", login)
    log.Println("Listening Siamese Fireback ...")
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}