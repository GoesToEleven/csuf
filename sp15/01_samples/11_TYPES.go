package main

import "fmt"

type mySentence string

func main() {

	var message mySentence = "Hello World!"

	// var message string = "Hello World!"

	fmt.Println(message)

}

// go is not an OOP language
// the type system that exists in go
// -- makes it so you don't need classes
// -- gives you more flexibility b/c you're not constrained by class requirements
// instead of using classes, we have user defined types
