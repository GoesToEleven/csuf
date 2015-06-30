package main

import (
	"log"
	"os"
	"html/template"
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
	tmpl, _ := template.ParseFiles("templates/file1.html", "templates/partial.html")
	err := tmpl.ExecuteTemplate(os.Stdout, "file1.html", u)
	if err != nil {
		log.Panic(err)
	}
}

/*
notes:
These files correspond to this training:
http://www.metacasts.tv/casts
*/
