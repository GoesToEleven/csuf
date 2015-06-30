package main

import "fmt"

// func counter(n int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(n, ":", i)
// 	}
// }

func main() {
	// counter(0)
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}

// this example has no concurency
// control flow is sequential, top to bottom
