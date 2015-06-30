package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 6; i++ {
		fmt.Println(i)
	}
	fmt.Println("And again...")
	for i := 1; ; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("Now the even numbers:")
	for i := 1; i < 6; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Now from a slice:")
	for _, current := range nums {
		fmt.Println(current)
	}
	mymap := map[string]string{
		"Corey":  "Dihel",
		"Taylor": "Dihel",
		"Alex":   "Taylor",
	}
	for key, value := range mymap {
		fmt.Println(key, value)
	}
}
