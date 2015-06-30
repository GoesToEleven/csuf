package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Emails []string
}

const tmpl = `The name is {{.Name}}.
{{range .Emails}}
    His email id is {{.}}
{{end}}
`

func main() {
	person := Person{
		Name:   "Satish",
		Emails: []string{"satish@rubylearning.org", "satishtalim@gmail.com"},
	}

	// STEP 1: create a new template
	t := template.New("Person template")

	// STEP 2: parse the string into the template
	// in lay terms: "give the template your form letter"
	// in lay terms: "put your form letter into the template"
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// STEP 3: execute the template
	//merge template with data
	err = t.Execute(os.Stdout, person)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}
