package main

import "fmt"

// TYPE
type mySentence string

// TYPES CAN CONTAIN METHODS
func (s mySentence) eatChocolate() {
	fmt.Println("METHOD: EAT MORE CHOCOLATE NOW")
}
func (s *mySentence) eatChocolate2() {
	s.eatChocolate()
	*s = "rawr"
}

func main() {

	// TYPES CAN CONTAIN DATA
	var message mySentence = "Hello World!"
	fmt.Println("DATA: " + message)

	// TYPES CAN USE METHODS
	message.eatChocolate()
	message.eatChocolate2()
	fmt.Println(message)

}
