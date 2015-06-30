package main

import (
	"fmt"
	"net/http"
)

func oneHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the oneHandler")
}

func twoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the twoHandler")
}

func threeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the threeHandler")
}

func main() {
	http.HandleFunc("/", oneHandler)
	http.HandleFunc("/two/", twoHandler)
	http.HandleFunc("/three/", threeHandler)
	http.ListenAndServe(":8080", nil)
}
