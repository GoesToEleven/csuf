package main

import "fmt"

// FOR
// -- clause
// ---- init; cond; post
// -- condition
// -- range

func main() {

	i := 0
	for i < 10 {
		if i >= 7 {
			break
		}
		fmt.Println(i)
		i++
	}

	// what we had before:

	// i := 0
	// for i < 10 {
	//   fmt.Println(i)
	//   i++
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
}
