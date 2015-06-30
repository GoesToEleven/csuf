package main

import (
	"fmt"
)

func main() {

	// Maps - Shorthand Notation
	myGreeting := map[string]string{
		"Tim":     "Good morning!",
		"Jenny":   "Bonjour!",
		"Medhi":   "Buenos dias!",
		"Marcus":  "Bongiorno!",
		"Julian":  "Ohayo!",
		"Sushant": "Selamat pagi!",
		"Jose":    "Gutten morgen!",
	}

	myGreeting["Me"] = "Hello!"

	fmt.Println(myGreeting["Me"])

	if val,exists := myGreeting["Tim"]; exists {
		fmt.Println(val)
	} else {
		fmt.Println("Doesn't exist")
	}

	delete(myGreeting, "Tim")


	if val,exists := myGreeting["Tim"]; exists {
		fmt.Println(val)
	} else {
		fmt.Println("Doesn't exist")
	}
}
