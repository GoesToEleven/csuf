package main

import "fmt"

// FOR
// -- clause
// ---- init; cond; post
// -- condition
// -- range

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, " - ", j)
		}
	}

	// will this, below, run?
	// fmt.Println(i)

}
