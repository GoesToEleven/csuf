package main

import "fmt"

type Contact struct {
	name     string
	greeting string
}

func Greet(contact Contact) {
	fmt.Println(contact.name)
	fmt.Println(contact.greeting)
}

func main() {

	var s = Contact{}
	s.name = "Marcus"
	s.greeting = "Hello!"
	Greet(s)

	var t = Contact{"Medhi", "Good to see you!"}
	Greet(t)

	u := Contact{"Sushant", "Glad you're in class!"}
	Greet(u)

	v := Contact{}
	v.name = "John"
	v.greeting = "We're learning great things!"
	Greet(v)

}
