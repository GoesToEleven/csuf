package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// THIS WAY
// func main() {
// 	fmt.Println(add(1, 2, 3))
// }

// OR THIS WAY
// We can also pass a slice of ints by following the slice with ...
func main() {
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))
}

/*
By using ... before the type name of the last parameter you can indicate
that it takes zero or more of those parameters. In this case we take zero
or more ints. We invoke the function like any other function except
we can pass as many ints as we want.

This is precisely how the fmt.Println function is implemented:

func Println(a ...interface{}) (n int, err error)

*/
