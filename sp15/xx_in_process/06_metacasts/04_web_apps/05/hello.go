package main

import (
	"fmt"
	"net/http"
)

func AnyName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)

}

func main() {
	http.HandleFunc("/", AnyName)

	http.ListenAndServe(":8080", nil)
}
