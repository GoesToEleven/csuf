package main

import "fmt"
import "math"
import m "github.com/goestoeleven/math"

func main() {

	myNumberCast := float64(36)
	sqrtOfNumber := math.Sqrt(myNumberCast)

	numberSlice := []int{5, 10, 15}
	avgOfNumbers := m.Average(numberSlice)

	fmt.Println(sqrtOfNumber)
	fmt.Println(avgOfNumbers)
}

// IF A WEIRD ERROR SHOWS - TWO PANELS OPEN, ONLY ONE PRINTLN SHOWS
// How can we isolate it and understand it, if not also solve it?
// answer down below ...

/*
notes
www.golang.org/pkg
www.golang-book.com/11/index.htm


































Answer:
runs correct from terminal; must be an Atom.io error
*/
