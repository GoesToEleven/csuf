/*
There is an inefficiency in our previous code:
renderTemplate calls ParseFiles every time a page is rendered.
A better approach would be to call ParseFiles once at program initialization,
parsing all templates into a single *Template.
Then we can use the ExecuteTemplate method to render a specific template.
https://golang.org/pkg/html/template/#Template.ExecuteTemplate
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

// First we create a global variable named templates, and initialize it with ParseFiles.

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

/*
The function template.Must is a convenience wrapper
that panics when passed a non-nil error value, and otherwise
returns the *Template unaltered. A panic is appropriate here;
if the templates can't be loaded the only sensible thing to do is exit the program.

The ParseFiles function takes any number of string arguments
that identify our template files, and parses those files into templates
that are named after the base file name. If we were to add more templates
to our program, we would add their names to the ParseFiles call's arguments.
*/

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

// modify the renderTemplate function to call the templates.ExecuteTemplate method
// with the name of the appropriate template:

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
WHAT WE HAD BEFORE
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
*/

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

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// defaultHandler ADD BY TODD MCLEOD
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='http://localhost:8080/view/testpage'>"+
		"http://localhost:8080/view/testpage</a>")
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
