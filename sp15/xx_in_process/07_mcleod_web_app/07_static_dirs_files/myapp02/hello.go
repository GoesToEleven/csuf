package myapp02

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("index.html"))

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
