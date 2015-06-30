package main

import (
	"html/template"
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
	u := User{"Mark", "Bates", "mark@example.com", 37}
	tmpl, _ := template.ParseFiles("templates/file4.tmpl", "templates/file5.tmpl")
	err := tmpl.ExecuteTemplate(os.Stdout, "file4.tmpl", u)
	if err != nil {
		log.Panic(err)
	}
}
