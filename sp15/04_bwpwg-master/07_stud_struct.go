package main

import (
	"log"
	"os"
	"text/template"
)

type Student struct {
	//exported field since it begins
	//with a capital letter
	Name string
}

func main() {
	//define an instance
	s := Student{"Medhi"}

	// STEP 1: create a new template
	// STEP 2: parse the string into the template
	// STEP 3: execute the template

	// STEP 1: create a new template with some name
	tmpl := template.New("test")

	// STEP 2: parse the string into the template
	// in lay terms: "give the template your form letter"
	// in lay terms: "put your form letter into the template"
	tmpl, err := tmpl.Parse("Hello {{.Name}}!")
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// STEP 3: execute the template
	//merge template 'tmpl' with content of 's'
	// lay terms: "merge your data into the form letter"
	err1 := tmpl.Execute(os.Stdout, s)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
		return
	}
}
