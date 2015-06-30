package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/pat"
)

func IndexUsers(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Users Index")
}

func ShowUser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Users Show: %s", req.URL.Query().Get(":id"))
}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Fprintf(res, "Users Create: %s", body)
}

func NewMux() http.Handler {
	pat := pat.New()

	pat.Get("/posts", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "POSTS!")
	})
	pat.Get("/users/{id}", ShowUser)
	pat.Get("/users", IndexUsers)
	pat.Post("/users", CreateUser)
	return pat
}

func main() {
	http.ListenAndServe(":3000", NewMux())
}
