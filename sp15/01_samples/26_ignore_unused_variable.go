package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func Greet(person Contact) {
	// this is what we had before:
	// fmt.Println(CreateMessage(contact.greeting, contact.name))

	greeting, _, _ := CreateMessage(person.greeting, person.name)

	// IF WE COMMENT OUT ONE OF THE BELOW WE GET AN ERROR
	// you can't declare a return and never use it - this causes an error
	// go eliminates warnings and just gives errors
	// generaly speaking, if you shouldn't do it, you can't do it
	// to ignore a returned variable, use the underscore
	fmt.Println(greeting)
}

func CreateMessage(greeting, name string) (string, string, int) {
	myNum := 120
	return greeting + " " + name, "\nHey, " + name + "\n", myNum
}

func main() {

	var t = Contact{"Good to see you,", "Medhi"}
	Greet(t)

	u := Contact{"Glad you're in class,", "Sushant"}
	Greet(u)

	v := Contact{}
	v.greeting = "We're learning great things,"
	v.name = "Max"
	Greet(v)
}
