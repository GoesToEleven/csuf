package main

import "fmt"

func main() {

	var customerNumber []int
	customerNumber = make([]int, 3)
	// 3 is length & capacity
	// // length - number of elements referred to by the slice
	// // capacity - number of elements in the underlying array
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15
	customerNumber[3] = 33

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])
	fmt.Println(customerNumber[3])

	// var greeting []string
	// greeting = make([]string, 3, 5)
	// // 3 is length - number of elements referred to by the slice
	// // 5 is capacity - number of elements in the underlying array
	// // you could also do it like this
	//

	// greeting[0] = "Good morning!"
	// greeting[1] = "Bonjour!"
	// greeting[2] = "dias!"
	// greeting[3] = "Bongiorno!"
	// greeting[4] = "Ohayo!"
	// greeting[5] = "Selamat pagi!"

	// fmt.Println(greeting[4])

}

/*
notes
https://docs.google.com/document/d/17AgfOseB1Pm_9lUxweCA47X-Gemaix3yx1IvoQMYYpI/edit#heading=h.658d4v7m9aao
*/
