package main

import (
	"fmt"
)

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	a, b := 1, 2
	fmt.Println("Original:", a, b)
	swap(&a, &b)
	fmt.Println("Swapped:", a, b)
}
