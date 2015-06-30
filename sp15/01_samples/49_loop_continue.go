package main

import "fmt"

// FOR
// -- clause
// ---- init; cond; post
// -- condition
// -- range

func main() {

	i := 0
	for i < 50 {

		if i >= 41 {
			break
		}

		if i%2 == 0 {
			i++
			continue
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
