package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Path:", r.URL.Path)
	fmt.Fprintln(w, "scheme:", r.URL.Scheme)
	fmt.Fprintln(w, "Opaque:", r.URL.Opaque)
	fmt.Fprintln(w, "User:", r.URL.User)
	fmt.Fprintln(w, "Host:", r.URL.Host)
	fmt.Fprintln(w, "RawQuery:", r.URL.RawQuery)
	fmt.Fprintln(w, "Fragment:", r.URL.Fragment)
	fmt.Fprintln(w, "String():", r.URL.String())
	// fmt.Fprintln(w, "ResolveReference():", r.URL.ResolveReference())
	fmt.Fprintln(w, "RequestURI():", r.URL.RequestURI())
	fmt.Fprintln(w, "Query():", r.URL.Query())
	fmt.Fprintln(w, "IsAbs():", r.URL.IsAbs())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

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
*/
