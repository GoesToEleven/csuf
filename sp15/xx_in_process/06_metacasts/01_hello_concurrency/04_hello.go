package main

import (
	"fmt"
)

func Greet(name string) {
	fmt.Println("Hello " + name)
}

func GreetNames(names []string, suffix string) {
	for _, n := range names {
		Greet(n + suffix)
	}
}

func main() {
	firstNames := []string{
		"Mark",
		"Rachel",
		"Dylan",
		"Leo",
	}

	comm := make(chan string)

	go func() {
		GreetNames(firstNames, " <C> ")
		comm <- "Finished greeting names concurrently ..."
	}()

	GreetNames(firstNames, " <M> ")

	fmt.Println(<-comm)

}
