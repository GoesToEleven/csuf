package main

import "fmt"

func main() {

	x := 0

	increment := func() int {
		x++
		return x
	}

	fmt.Println(x)
	fmt.Println(increment())
	fmt.Println(increment())

}

/*

CLOSURE
A function like this together with the non-local variables it references
is known as a closure. In this case increment and the variable x form the closure.

*/
