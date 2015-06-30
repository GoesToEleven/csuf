package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	u := User{"Todd", "McLeod", "tuddleymc@gmail.com", 44}
	body, _ := ioutil.ReadFile("templates/file1.html")

	tmpl, err := template.New("AnyName").Parse(string(body))
	if err != nil {
		log.Panic(err)
	}

	tmpl.Execute(os.Stdout, u)
}
