package main

import "fmt"

func main() {
	const x string = "Hello World"
	fmt.Println(x)

	// the code below won't work
	// x = "some other string"
	// fmt.Println(x)
}

/*
Constants
variables whose values cannot be changed
created in the same way you create variables but instead of using the var keyword
we use the const keyword
*/
