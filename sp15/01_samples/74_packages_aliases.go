package main

import "fmt"
import m "github.com/goestoeleven/math"

func main() {

	numberSlice := []int{5, 10, 15}

	avgOfNumbers := m.AverageCorrect(numberSlice)

	fmt.Println(avgOfNumbers)

}

/*
notes
www.golang.org/pkg
www.golang-book.com/11/index.htm
*/
