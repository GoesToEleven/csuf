package main

import (
	"fmt"
)

func print_fav(name string, num int) {
	fmt.Println(name, "likes", num)
}

func main() {
	fav_nums := map[string]int{
		"Corey":   17,
		"Allen":   21,
		"Antonio": 13,
		"Michael": 1,
		"Alex":    3,
	}
	for key, value := range fav_nums {
		print_fav(key, value)
	}
}
