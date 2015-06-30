package main

import "fmt"

// RECURSION
// a function which calls itself

func factorial(x uint64) uint64 {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(65))
}

/*
Closure and recursion are powerful programming techniques which form the basis
of a paradigm known as functional programming.
*/
