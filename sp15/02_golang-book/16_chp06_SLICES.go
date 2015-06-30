package main

import "fmt"

func main() {

	// ARRAY
	var myArray [5]int
	myArray[4] = 100
	fmt.Println(myArray)
	fmt.Println("***********")

	// ARRAY SHORTHAND
	myArrayMas := [5]float64{98, 93, 77, 82, 83}
	fmt.Println(myArrayMas)
	fmt.Println("***********")

	// SLICE shorthand
	// The only difference between this and an array
	// is the missing length between the brackets.
	mySlice := []int{1, 5, 15, 20, 25, 30}
	fmt.Println(mySlice)
	fmt.Println("***********")

	// SLICE LONG WAY
	var customerNumber []int
	customerNumber = make([]int, 3)
	// 3 is length & capacity
	// // length - number of elements referred to by the slice
	// // capacity - number of elements in the underlying array
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15
	fmt.Println(customerNumber)
	// fmt.Println(customerNumber[0])
	// fmt.Println(customerNumber[1])
	// fmt.Println(customerNumber[2])

	// var greeting []string
	// greeting = make([]string, 3, 5)
	// // 3 is length - number of elements referred to by the slice
	// // 5 is capacity - number of elements in the underlying array

	// LOOP OVER RANGE
	// for i, currentEntry := range mySlice {
	// 	fmt.Println(i, " - ", currentEntry)
	// }
}
