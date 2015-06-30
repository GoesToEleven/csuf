/*
CODE INJECTION
http://en.wikipedia.org/wiki/Code_injection

As you may have observed, this program has a serious security flaw:
a user can supply an arbitrary path to be read/written on the server.
To mitigate this, we can write a function to validate the title with a regular expression.

add "regexp" to the import list
*/

package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

/*
The function regexp.MustCompile will parse and compile the regular expression,
and return a regexp.Regexp. MustCompile is distinct from Compile in that it will panic
if the expression compilation fails, while Compile returns an error as a second parameter.

write a function getTitle
that uses the validPath expression to validate path and extract the page title:
*/

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

/*
If the title is valid, it will be returned along with a nil error value.
If the title is invalid, the function will write a "404 Not Found" error
to the HTTP connection, and return an error to the handler.
To create a new error, we have to import the errors package.
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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

/*
WHAT WE HAD BEFORE
	title := r.URL.Path[len("/edit/"):]
*/

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

/*
WHAT WE HAD BEFORE
	title := r.URL.Path[len("/view/"):]
*/

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

/*
WHAT WE HAD BEFORE
  title := r.URL.Path[len("/save/"):]
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
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
