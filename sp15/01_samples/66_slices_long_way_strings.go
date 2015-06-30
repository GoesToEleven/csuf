package main

import "fmt"

func main() {

	var customerName []string
	customerName = make([]string, 3, 5)
	// 3 is length, 5 is capacity
	// // length - number of elements referred to by the slice
	// // capacity - number of elements in the underlying array
	customerName[0] = "Rikhil"
	customerName[1] = "Akashdeep"
	customerName[2] = "Ishan"
	// customerName[3] = "Tim"

	fmt.Println(customerName[0])
	fmt.Println(customerName[1])
	fmt.Println(customerName[2])
	// fmt.Println(customerName[3])
}

/*
notes
https://docs.google.com/document/d/17AgfOseB1Pm_9lUxweCA47X-Gemaix3yx1IvoQMYYpI/edit#heading=h.658d4v7m9aao
*/
