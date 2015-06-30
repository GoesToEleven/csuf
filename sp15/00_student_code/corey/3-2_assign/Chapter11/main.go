package main

import (
	"fmt"

	"github.com/Saxleader/golang/3-2_assign/Chapter11/min_max"
)

func main() {
	list := []float64{25, 78, 61, 111, 2, 45, 69, 71, 3}
	fmt.Println("The list is:", list)
	fmt.Println("The min value is:", min_max.Min(list))
	fmt.Println("The max value is:", min_max.Max(list))
}
