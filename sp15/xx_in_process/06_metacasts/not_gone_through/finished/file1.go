package main

import (
	"io/ioutil"
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
	body, _ := ioutil.ReadFile("templates/file1.tmpl")
	tmpl, err := template.New("Some Name").Parse(string(body))
	if err != nil {
		log.Panic(err)
	}
	tmpl.Execute(os.Stdout, u)
}
