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
	u := User{"Mark", "Bates", "mark@example.com", 37}
	tmpl, err := template.New("Some Name").Parse("Name: {{.FirstName}} {{.LastName}}\nEmail: {{.Email}}\nAge: {{.Age}}")
	if err != nil {
		log.Panic(err)
	}
	tmpl.Execute(os.Stdout, u)
}
