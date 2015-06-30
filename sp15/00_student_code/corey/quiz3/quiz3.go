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
	fruit := pies[:4]
	fmt.Println("\nFruit Pies:")
	for _, current := range fruit {
		fmt.Println(current)
	}
	my_str := "Corey Dihel"
	fmt.Println("\nFull Name:", my_str)
	fmt.Println("First Name:", my_str[:5])
	fmt.Println("Last Name:", my_str[6:])
}
