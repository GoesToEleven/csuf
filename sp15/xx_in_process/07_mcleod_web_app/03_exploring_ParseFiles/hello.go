package hello

import (
	"html/template"
	"net/http"
)

var (
	// OLD CODE:
	// guestbookForm []byte
	// signTemplate  = template.Must(template.ParseFiles("templates/guestbook.html"))
	tmpl, _ = template.ParseFiles("templates/guestbook.html", "templates/guestbookform.html")
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "guestbookform.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sign(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "guestbook.html", r.FormValue("content"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
NOTES:
http://golang.org/pkg/html/template/#ParseFiles
http://golang.org/pkg/html/template/#Template.ExecuteTemplate
http://golang.org/pkg/net/http/#Error
http://blog.golang.org/error-handling-and-go
*/
