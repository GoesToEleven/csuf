package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

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
