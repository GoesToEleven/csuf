package main

import "fmt"

// FUNCTIONS INSIDE FUNCTIONS
// It is possible to create functions inside of functions

func main() {

	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 1))

}

// assigning a function to a variable also known as function literal
