package main

import (
  "net/http"
  "text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content Type", "text/html")
			templates := template.New("template")
			templates.New("test").Parse(doc)
			templates.New("header").Parse(header)
			templates.New("footer").Parse(footer)
			context := Context{
				[3]string{"Lemon", "Orange", "Apple"},
				"the title",
			}
			templates.Lookup("test").Execute(w, context)
		})
	
	http.ListenAndServe(":8000", nil)
}

const doc = `
{{template "header" .Title}}
  <body>
    <h1>List of Fruit</h1>
    <ul>
      {{range .Fruit}}
      	<li>{{.}}</li>
      {{end}}
    </ul>
  </body>
{{template "footer"}}
`

const header = `
<!DOCTYPE html>
<html>
  <head><title>{{.}}</title></head>
`

const footer = `
</html>
`

type Context struct {
	Fruit [3]string
	Title string
}