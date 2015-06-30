package main

import "fmt"

func main() {
	var y [5]int
	y[4] = 100
	fmt.Println(y)

	fmt.Println("***********")

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

	// you can also use RANGE for loop
	var totalMas float64 = 0
	for _, value := range x {
		totalMas += value
	}
	fmt.Println(totalMas / float64(len(x)))
	// QUESTION: why do we need to cast the len(x) above?
	// particularly when we were able to divide a float by an int
	// in 07_chap03_types_floating_point.go
	// (note: I don't have the answer)

	// ARRAY SHORTHAND
	z := [5]float64{98, 93, 77, 82, 83}
	for i, value := range z {
		fmt.Println(value, " - ", z[i])
	}

}
