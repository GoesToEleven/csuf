package main

import (
	"fmt"
)

func main() {
	names := []string{
		"Mark",
		"Rachel",
		"Dylan",
		"Leo",
	}

	for _, name := range names {
		fmt.Println("Hello " + name)
	}
}
