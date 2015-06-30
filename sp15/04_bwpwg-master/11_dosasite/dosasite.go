package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// http://golang.org/pkg/net/http/#StripPrefix
	// To serve a directory on disk (/public) under an alternate URL
	// path (/tmpfiles/), use StripPrefix to modify the request
	// URL's path before the FileServer sees it:
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", fs))

	fmt.Println("Listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

/*
handle
handlefunc
handler
handlerfunc

http://golang.org/pkg/net/http/#Handle				// takes a string & a handler
http://golang.org/pkg/net/http/#HandleFunc		// takes a string & a handlerfunc
http://golang.org/pkg/net/http/#Handler				// defines the handler interface
http://golang.org/pkg/net/http/#HandlerFunc		// a func that implements the handler interface

*/
