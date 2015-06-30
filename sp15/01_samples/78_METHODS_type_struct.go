package main

import "fmt"

// STRUCT TYPE
type contact struct {
	name     string
	greeting string
}

// STRUCTS (TYPES) CAN CONTAIN METHODS
func (c contact) sayHello() {
	fmt.Println("METHOD: " + c.name + " says hello.")
}

func main() {
	// STRUCTS (TYPES) CAN CONTAIN DATA
	var friend = contact{"Marcus", "Hello!"}
	fmt.Println("DATA: " + friend.name)
	fmt.Println("DATA: " + friend.greeting)

	// STRUCTS (TYPES) CAN USE METHODS
	friend.sayHello()

}

// we can use a STRUCT like a class

// go is not an OOP language
// the type system that exists in go
// -- makes it so you don't need classes
// -- gives you more flexibility b/c you're not constrained by class requirements
// instead of using classes, we have user defined types

// NOTE: in the method, (c Contact) is called a "receiver"
// it isn't called a "parameter"

/*















	var c = contact{"Marcus", "Hello!"}

	fmt.Println(c.name)
	fmt.Println(c.greeting)

*/
