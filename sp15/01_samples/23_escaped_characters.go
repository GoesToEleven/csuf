package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func Greet(contact Contact) {
	fmt.Println(CreateMessage(contact.greeting, contact.name))
}

func CreateMessage(greeting, name string) (string, string) {
	return greeting + " " + name, "\nHey, " + name + "\n\n"
}

func main() {

	var t = Contact{"Good to see you,", "Medhi"}
	Greet(t)

	u := Contact{"Glad you're in class,", "Sushant"}
	Greet(u)

	v := Contact{}
	v.greeting = "We're learning great things,"
	v.name = "Marcus"
	Greet(v)
}

// put parentheses around return types and comma separate them
// use comma to separate the multiple returns
