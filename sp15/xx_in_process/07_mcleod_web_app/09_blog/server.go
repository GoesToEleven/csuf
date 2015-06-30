package server

import (
	"appengine"
	"appengine/user"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"site/includes/document_top.html",
	"site/includes/document_bottom.html",
	"site/index.html"))

func init() {
	http.HandleFunc("/", root)
	//	 http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	u := user.Current(c)
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

	err := tmpl.ExecuteTemplate(w, "index.html", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func sign(w http.ResponseWriter, r *http.Request) {
// 	err := tmpl.ExecuteTemplate(w, "guestbook.html", r.FormValue("content"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

/*
NOTES:
http://golang.org/pkg/html/template/#Must
https://golang.org/doc/effective_go.html#data
*/
