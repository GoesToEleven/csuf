package main

import (
	"fmt"
	"log"
	"net/http"
)

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!!")
}

func main() {
	http.Handle("/", &HomePageHandler{})

	barHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Bar!")
	}
	http.HandleFunc("/bar", barHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
