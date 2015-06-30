package hello

import (
    "fmt"
    "net/http"

    "appengine"
    "appengine/user"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    // [START new_context]
    c := appengine.NewContext(r)
    // [END new_context]
    // [START get_current_user]
    u := user.Current(c)
    // [END get_current_user]
    // [START if_user]
    if u == nil {
        url, err := user.LoginURL(c, r.URL.String())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Location", url)
        w.WriteHeader(http.StatusFound)
        return
    }
    	w.Header().Set("myLocation", r.URL.String())
        w.Header().Set("myname", "todd")
        w.WriteHeader(http.StatusOK)
    // [END if_user]
    // [START output]
    fmt.Fprintf(w, "Hello, %v!", u)
    // [END output]
}