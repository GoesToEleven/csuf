package main

import (
	"fmt"
)

func main() {
	
	//for loop
	for i:= 0; i < 10; i++ {
		fmt.Println(i)
	}
	
	//continue and break
	for i:=0; i < 20; i++ {
		if i % 2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
		
		if i > 11 {
			break
		}
	}
	
	//slices
	mySlice := []int{1,5,15,20,25,30}
	
	for i, currentEntry := range mySlice {
		fmt.Println(i, " - ", currentEntry)
	}
	fmt.Print("[2:]")
	fmt.Println(mySlice[2:])


	//string slicing
	myString := "don't cut me!"
	fmt.Println(myString)
	fmt.Println(myString[6:])

	//more string slicing
	myStringSlice := []string{"one", "two", "three", "four", "five"}
	fmt.Println(myStringSlice[1:][1][1:])

	//loops?
	for key, val := range(myStringSlice) {
		fmt.Println(key, " - ", val)
	}
}
