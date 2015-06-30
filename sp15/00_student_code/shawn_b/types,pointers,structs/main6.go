package main

import (
	"fmt"
)

func main() {

	var customerName []string
	customerName = make([]string, 3, 5)
	// 3 is length, 5 is capacity
	// // length - number of elements referred to by the slice
	// // capacity - number of elements in the underlying array
	customerName[0] = "Rikhil"
	customerName[1] = "Akashdeep"
	customerName[2] = "Ishan"
	customerName = append(customerName, "Tim")
	customerName = append(customerName, "Tim2")
	customerName = append(customerName, "Tim3")
	customerName = append(customerName, "Tim4")

	fmt.Println(customerName)
	//fmt.Println(customerName[1])
	//fmt.Println(customerName[2])

	mySlice := []int{1,2,3,4}
	myOtherSlice := []int{5,6,7,8,9}

	var myBiggerSlice []int
	myBiggerSlice = append(mySlice, myOtherSlice...)

	fmt.Println("myslice - ", mySlice)
	fmt.Println("myOtherSlice - ", myOtherSlice)
	fmt.Println("myBiggerSlice - ", myBiggerSlice)

}
