package main

import (
	"fmt"
	"net/http"
)

// The function handler is of the type http.HandlerFunc
// It takes an http.ResponseWriter and http.Request as its arguments
func handler(w http.ResponseWriter, r *http.Request) {
	// An http.ResponseWriter value assembles the HTTP server's response
	// by writing to it, we send data to the HTTP client
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
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
printf % - see:
https://gobyexample.com/string-formatting
gowiki/04_string_formatting.go

An http.Request is a data structure that represents the client HTTP request.
r.URL.Path is the path component of the request URL.
The trailing [1:] means "create a sub-slice of Path from the 1st character to the end."
This drops the leading "/" from the path name.


OTHER INFO


HANDLEFUNC
http.HandleFunc
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
The documentation for ServeMux explains how patterns are matched.
http://golang.org/pkg/net/http/#HandleFunc

Processing HTTP requests with Go is primarily about two things: ServeMuxes and Handlers.
http://www.alexedwards.net/blog/a-recap-of-request-handling

SERVEMUX
A ServeMux is essentially a HTTP request router (or multiplexor)
It compares incoming requests against a list of predefined URL paths,
and calls the associated handler for the path whenever a match is found.
-- ServeMux
-- -- http://golang.org/pkg/net/http/#ServeMux

HANDLERS
Handlers are responsible for writing response headers and bodies.
Almost any object can be a handler, so long as it satisfies the http.Handler interface.
That means it must have a ServeHTTP method with the following signature:
ServeHTTP(http.ResponseWriter, *http.Request)

LISTENANDSERVE
http://golang.org/pkg/net/http/#ListenAndServe
func ListenAndServe(addr string, handler Handler) error
ListenAndServe listens on the TCP network address addr and then calls
Serve with handler to handle requests on incoming connections.
Handler is typically nil, in which case the DefaultServeMux is used.

SERVEMUX
see above

*/
