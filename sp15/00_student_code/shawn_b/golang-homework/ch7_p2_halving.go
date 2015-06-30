/*
Write a function which takes an integer and halves it and returns true if it was even or false if it was odd.
For example half(1) should return (0, false) and half(2) should return (1, true).
 */

package main

import (
	"fmt"
)

func half(x int) (int, bool) {
	if x % 2 == 0 {
		return x/2, true
	} else {
		return x/2, false
	}
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}
