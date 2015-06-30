package main

import "fmt"

func main() {

	message := "Hello World!"

	var greeting *string = &message

    var g2 string = *greeting

    var g3 *string = &g2

	fmt.Println("message - string ",message)
	fmt.Println("greeting - memory address ", greeting)
	fmt.Println("g2 - string", g2)
	fmt.Println("g3 - memory address", g3)

	// the * says, "Show me what you're pointing towards."
}
