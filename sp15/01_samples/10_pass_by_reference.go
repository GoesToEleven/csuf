package main

import "fmt"

func main() {

	// set message to a string literal
	message := "Hello World!"
	fmt.Println("message: ", message)

	// set greeting to point towards memory location of message
	var greeting *string = &message
	fmt.Println("greeting shows memory location of message: ", greeting)

	// greeting is the memory location of message; change string literal to which it points
	*greeting = "hi"
	// fmt.Println(*greeting)

	fmt.Println("message and greeting: ", message, *greeting)
}
