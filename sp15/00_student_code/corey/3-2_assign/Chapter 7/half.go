package main

import (
	"fmt"
)

func half(num int) (int, bool) {
	h := num / 2
	if num%2 == 0 {
		return h, true
	}
	return h, false
}

func main() {
	a, b := 1, 2
	c, d := half(a)
	e, f := half(b)
	fmt.Println(a, c, d)
	fmt.Println(b, e, f)
}
