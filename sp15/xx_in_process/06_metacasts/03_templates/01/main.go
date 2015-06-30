package main

import (
	"log"
	"os"
	"text/template"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}

func main() {
	u := User{"Todd", "McLeod", "tuddleymc@gmail.com", 44}
	tmpl, err := template.New("AnyName").Parse("Name: {{.FirstName}} {{.LastName}}\nEmail: {{.Email}}\nAge: {{.Age}}")
	if err != nil {
		log.Panic(err)
	}

	tmpl.Execute(os.Stdout, u)
}
