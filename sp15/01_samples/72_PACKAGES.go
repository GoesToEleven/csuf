package main

import "fmt"
import "github.com/goestoeleven/math"

// notice the import path begins in the folder AFTER the src folder

func main() {

	numberSlice := []int{5, 10, 15}

	sumOfNumbers := math.Sum(numberSlice)

	fmt.Println(sumOfNumbers)
}

/*
notes
www.golang-book.com/11/index.htm
*/
