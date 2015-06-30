package main

import "fmt"

func main() {

	message := "Hello World!"

	var greeting *string = &message

	fmt.Println(message, greeting)

	// the * makes a pointer of a certain type
	// the above code makes greeting a pointer to a string
	// the & before a variable says, "Give me the memory location of that variable"
}
