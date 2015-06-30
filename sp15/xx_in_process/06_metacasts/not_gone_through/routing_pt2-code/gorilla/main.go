package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func IndexUsers(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Users Index")
}

func ShowUser(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Fprintf(res, "Users Show: %s", vars["id"])
}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Fprintf(res, "Users Create: %s", body)
}

func IndexPosts(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "POSTS!")
}

func NewMux() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/posts", IndexPosts)
	r.HandleFunc("/users/{id:[0-9]+}", ShowUser).Methods("GET")
	r.HandleFunc("/users", IndexUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	return r
}

func main() {
	http.ListenAndServe(":3000", NewMux())
}
