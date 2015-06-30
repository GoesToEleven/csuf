package server

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"site/includes/document_top.html",
	"site/includes/document_bottom.html",
	"site/index.html"))

func init() {
	http.HandleFunc("/", root)
	// http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}