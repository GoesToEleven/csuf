package main

import (
    "fmt"
    "net/http"
    "time"

    "appengine"
    "appengine/datastore"
)

type Employee struct {
    Name     string
    Role     string
    HireDate time.Time
    email    string
}

func init() {
    http.HandleFunc("/", myHandler)
    http.HandleFunc("/read", myRead)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    e1 := Employee{
        Name:     "James Dean",
        Role:     "Manager",
        HireDate: time.Now(),
        email:    "james@starship.com",
    }

    userKey := datastore.NewKey(c, "employee", e1.email, 0, nil)
    key, err := datastore.Put(c, userKey, &e1)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var e2 Employee
    if err = datastore.Get(c, key, &e2); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Wrote (Put) then read (Get) %q", e2.Name)
}

func myRead(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    var e1 Employee

    userKey := datastore.NewKey(c, "employee", "david@starship.com", 0, nil)

    if err = datastore.Get(c, userKey, &e1); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Retrieved %q", e1)
}
