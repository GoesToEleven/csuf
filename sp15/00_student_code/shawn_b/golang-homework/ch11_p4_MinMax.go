/*
We copied the average function from chapter 7 to our new package.
Create Min and Max functions which find the minimum and maximum values in a slice of float64s.
 */

package main

import(
	"fmt"
	"CSCI130-Go/golang-homework/golang_math"
)

func main() {
	xs := []float64{1,2, -1, 20,3,4}
	avg := math.Average(xs)
	fmt.Println(avg)
	fmt.Println(math.Min(xs))
	fmt.Println(math.Max(xs))
}
