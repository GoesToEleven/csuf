package main

import "fmt"

func f() (int, int) {
	return 5, 6
}

func main() {
	x, y := f()
	fmt.Println(x, y)
}

/*
Multiple values are often used to return an error value along with the result
(x, err := f()), or a boolean to indicate success (x, ok := f())
*/
