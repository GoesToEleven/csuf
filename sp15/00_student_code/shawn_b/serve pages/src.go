package hello

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/sign", sign)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func sign(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "No Autographs!")
}
