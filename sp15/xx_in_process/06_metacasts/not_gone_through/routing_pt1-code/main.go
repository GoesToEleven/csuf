package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	users := func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "" {
			switch req.Method {
			case "GET":
				fmt.Fprint(res, "Users Index")
			case "POST":
				body, _ := ioutil.ReadAll(req.Body)
				fmt.Fprintf(res, "Users Create: %s", body)
			}
		} else {
			fmt.Fprintf(res, "Users Show: %s", req.URL.Path)
		}
	}
	mux.HandleFunc("/posts", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "POSTS!")
	})
	mux.Handle("/users/", http.StripPrefix("/users/", http.HandlerFunc(users)))
	return mux
}

func main() {
	http.ListenAndServe(":3000", NewMux())
}
