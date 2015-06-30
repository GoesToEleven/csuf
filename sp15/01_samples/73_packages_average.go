package main

import "fmt"
import "github.com/goestoeleven/math"

func main() {

	numberSlice := []int{5, 10, 15}

	avgOfNumbers := math.Average(numberSlice)

	fmt.Println(avgOfNumbers)
}

// look at this code
// QUESTION: What is the average?

// notice the import path begins in the folder AFTER the src folder
// you might need to run GO INSTALL on the math dir from the terminal
// you might need to run GO INSTALL *every* *time* you change a package

/*
notes
www.golang-book.com/11/index.htm
*/
