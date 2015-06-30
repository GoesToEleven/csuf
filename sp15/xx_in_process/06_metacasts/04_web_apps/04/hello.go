package main

import (
	"fmt"
	"net/http"
)

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	http.Handle("/", &HomePageHandler{})

	http.ListenAndServe(":8080", nil)
}
