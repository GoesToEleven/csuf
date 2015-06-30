package main

import (
	"fmt"
	"net/http"
)

/*
NOTE:
at terminal "ctrl + c" exits web server
*/

// The function handler is of the type http.HandlerFunc
// It takes an http.ResponseWriter and http.Request as its arguments
func handler(w http.ResponseWriter, r *http.Request) {
	// An http.ResponseWriter value assembles the HTTP server's response
	// by writing to it, we send data to the HTTP client
	fmt.Fprintf(w, "Hi there.")
}

func main() {
	http.HandleFunc("/", handler)
	//http.HandleFunc tells the http package to handle all requests to the web root ("/") with handler
	http.ListenAndServe(":8080", nil)
	//http.ListenAndServe specifies that it should listen on port 8080 on any interface (":8080")
	//(Don't worry about its second parameter, nil, for now.)
	//This function will block until the program is terminated
}

/*
Print, Fprint, Println don't print to browser, only to the terminal
you need to have a writer (io.Writer) as a param which means using Fprint, Fprintf, Fprintln
https://golang.org/pkg/fmt/#Fprint

*/
