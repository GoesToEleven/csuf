package main

import (
	"fmt"
)

type str_test string

func (s str_test) testPrint(t str_test) {
	fmt.Println(s, t)
}

// printSlice is a function to print all of the elements in slice.
func printSlice(slice []int) {
	for _, current := range slice {
		fmt.Println(current)
	}
	fmt.Println()
}

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// printSlice(nums)
	// nums = append(nums, 6, 7)
	// printSlice(nums)
	var myArr [6]int
	other := myArr[:5]
	for key, _ := range other {
		other[key] = key + 1
	}
	fmt.Println("Length:", len(myArr), "\nCapacity:", cap(other))
	printSlice(other)
	other = append(other, 6)
	other = append(other, 7, 8)
	printSlice(other)
	for i := range myArr {
		fmt.Println(myArr[i])
	}
	var my_str str_test = "hello"
	my_str.testPrint(my_str)
}
