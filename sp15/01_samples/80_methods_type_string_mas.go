package main

import "fmt"

// TYPE
type mySentence string

// TYPES CAN CONTAIN METHODS
func (s mySentence) eatChocolate() {
	fmt.Println("METHOD: EAT MORE CHOCOLATE NOW", s)
}

func main() {

	// TYPES CAN CONTAIN DATA
	var message mySentence = "Hello World!"
	var justSaying mySentence = "Uhh, what drought?"
	fmt.Println("DATA: " + message)
	fmt.Println("DATA: " + justSaying)

	// TYPES CAN USE METHODS
	message.eatChocolate()
	justSaying.eatChocolate()

}

// Question:
// Can you access the data
// of each variable of type mySentence
// in the eatChocolate method?
