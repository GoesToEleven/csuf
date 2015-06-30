package main

import (
	"fmt"
)

func main() {
	pies := []string{"Apple", "Peach", "Cherry", "Blueberry", "Chocolate", "Pecan"}
	fmt.Println("Available Flavors of Pie:")
	for _, current := range pies {
		fmt.Println(current)
	}
}
