package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u)
}

// OUTPUT: https://google.com/search?q=golang

// reference:
// http://golang.org/pkg/net/url/#URL

/*
type URL struct {
        Scheme   string
        Opaque   string    // encoded opaque data
        User     *Userinfo // username and password information
        Host     string    // host or host:port
        Path     string
        RawQuery string // encoded query values, without '?'
        Fragment string // fragment for references, without '#'
}
A URL represents a parsed URL (technically, a URI reference).
The general form represented is:

scheme://[userinfo@]host/path[?query][#fragment]

URLs that do not start with a slash after the scheme are interpreted as:

scheme:opaque[?query][#fragment]

Note that the Path field is stored in decoded form:
/%47%6f%2f becomes /Go/
A consequence is that it is impossible to tell which slashes in the Path
were slashes in the raw URL and which were %2f. This distinction is rarely
important, but when it is a client must use other routines to parse the
raw URL or construct the parsed URL. For example, an HTTP server can consult
req.RequestURI, and an HTTP client can use
URL{Host: "example.com", Opaque: "//example.com/Go%2f"}
instead of
URL{Host: "example.com", Path: "/Go/"}.
*/
