package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    "time"
    "crypto/md5"
    "io"
    "strconv"
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

func valSelect(r *http.Request) bool {

    slice:=[]string{"apple","pear","banana"}

    for _, v := range slice {
        if v == r.Form.Get("fruit") {
            return true
        }
    }
    return false
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) // get request method
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, token)
    } else {
        // log in request
        r.ParseForm()

        validFruit := valSelect(r)

        if !validFruit {
            fmt.Fprintf(w, "that is not fruit")
        }

        token := r.Form.Get("token")
        if token != "" {
            // check token validity
        } else {
            // give error if no token
        }

        // logic part of log in
        fmt.Println("r.form:", r.Form)
        fmt.Println("fruit:", r.Form["fruit"])
        fmt.Fprintf(w, "fruit: %s", r.Form["fruit"])
    }
}

func main() {
    http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/login", login)
    log.Println("Listening Prince Ruspoli’s Turaco ...")
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}