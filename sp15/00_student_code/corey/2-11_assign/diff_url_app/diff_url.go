package diff

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/first", first)
	http.HandleFunc("/second", second)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Hello and welcome to the basic page of my app.\n")
	fmt.Fprint(rw, "Please check out the other pages: /first & /second")
}

func first(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Welcome to a different url of my app.")
}

func second(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Ok, by now, you can see that these different urls have different functions.")
}
