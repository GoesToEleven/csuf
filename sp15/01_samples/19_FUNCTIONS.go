package main

import "fmt"

// a bundle of variables of different types
// a variable of a variable
// you wrapped up a bunch of variables
type Contact struct {
	name     string
	greeting string
	age      int
}

func Greet(person Contact) {
	fmt.Println(person.name)
	fmt.Println(person.greeting)
}

func main() {
	// you can write this either way:

	// var c = Contact{
	// 	name:     "Marcus",
	// 	greeting: "Hello!",
	// }

	var c = Contact{}
	c.name = "Marcus"
	c.greeting = "Hello!"
	Greet(c)
}

// FUNCTIONS IN GO
// -- multiple return values
// -- use like any other type
// ---- similar to JavaScript
// ------ pass into other functions
// ------ declare them as variables
// ------ return them from functions
// -- function literals
// ---- you can declare functions inside other functions
