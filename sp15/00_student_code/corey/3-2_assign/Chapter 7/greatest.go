package main

import (
	"fmt"
)

func greatest(list ...int) (greatest int) {
	greatest = list[0]
	for _, value := range list {
		if value > greatest {
			greatest = value
		}
	}
	return
}

func main() {
	g := greatest(3, 6, 9, 4, 3, 6, 9, 54, 2, 67, 1)
	fmt.Println("The greatest value in the list is ", g)
}
