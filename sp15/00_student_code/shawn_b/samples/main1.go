package main

import (
	"fmt"
	tcm "CSCI130-Go/tmmath"
)

func newSum (intArrayParam [3]int) int {
	total := 0

	for _, x := range intArrayParam {
		total += x
	}

	return total
}

type places struct {
	name	mySent
	language	string
}

func (p places) sayHello() {
	fmt.Println(p.name, " speaks ", p.language)
}

type mySent string
func (s mySent) hoorah() {
	fmt.Println("METHOD: hoorah")
}

func (s mySent) retLang() string {
	return "I Speak ???"
}


func main() {

	//
	//mathy package things
	//
	numberSlice := []int{10,20,30}
	sumOfNumbers := tcm.Sum(numberSlice)
	fmt.Println(sumOfNumbers)
	avgOfNumbers := tcm.Average(numberSlice)
	fmt.Println(avgOfNumbers)

	//
	//slices
	//
	mySlice := []int{1,2,3,4}
	myOtherSlice := []int{5,6,7,8,9}

	var myBiggerSlice []int
	myBiggerSlice = append(mySlice, myOtherSlice...)

	fmt.Println("myslice - ", mySlice)
	fmt.Println("myOtherSlice - ", myOtherSlice)
	fmt.Println("myBiggerSlice - ", myBiggerSlice)

	var smallerThanBiggerSlice []int
	smallerThanBiggerSlice = append(mySlice[2:4], myOtherSlice[1:3]...)
	fmt.Println("smallerThanBiggerSlice - ", smallerThanBiggerSlice)

	//
	//slices vs arrays
	//
	numberSlice = []int{5,10,15}
	numberArray := [3]int{5,10,15}
	fmt.Println("slice sum = ", tcm.Sum(numberSlice))

	//doesn't work
	//fmt.Println("array sum = ", tcm.Sum(numberArray))

	fmt.Println("array slice sum = ", tcm.Sum(numberArray[:]))
	fmt.Println("array newSum - ", newSum(numberArray))

	//doesn't work
	//fmt.Println("slice newSum - ", newSum(numberSlice))
	fmt.Println("slice newSum - ", newSum([3]int{numberSlice[0], numberSlice[1], numberSlice[2]}))

	//
	// struct and method
	//
	var USA = places{
		name: "USA",
		language: "English",
	}
	USA.sayHello()

	//
	// string and method
	//
	USA.name.hoorah()

	//
	// string and method with return
	//
	fmt.Println(USA.name.retLang())
}
