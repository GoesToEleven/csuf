/*
HTTP
http://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol

HTTP STATUS CODES
http://en.wikipedia.org/wiki/List_of_HTTP_status_codes

HTTP/2
http://en.wikipedia.org/wiki/HTTP/2
*/

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

/*
http.Redirect function adds an HTTP status code of http.StatusFound (302)
and a Location header to the HTTP response
*/

// defaultHandler ADD BY TODD MCLEOD
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='http://localhost:8080/view/testpage'>"+
		"http://localhost:8080/view/testpage</a>")
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
