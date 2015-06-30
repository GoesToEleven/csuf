package main

import "fmt"

func main() {
	var message string
	message = "Hello World."

	var name string
	name = "Medhi"

	fmt.Println(message)

	// comment this out to see an error:
	// name declared and not used
	fmt.Println(name)
}
