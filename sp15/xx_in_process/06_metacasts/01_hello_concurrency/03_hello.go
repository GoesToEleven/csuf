package main

import (
	"fmt"
)

func Greet(name string) {
	fmt.Println("Hello " + name)
}

func GreetNames(names []string) {
	for _, n := range names {
		Greet(n)
	}
}

func main() {
	firstNames := []string{
		"Mark",
		"Rachel",
		"Dylan",
		"Leo",
	}
	GreetNames(firstNames)
}
