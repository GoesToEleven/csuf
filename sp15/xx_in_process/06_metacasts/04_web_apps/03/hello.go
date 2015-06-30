package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/bar", barWhatever)

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func barWhatever(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Bar!")
}
