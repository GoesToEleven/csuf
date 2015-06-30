package main

import "fmt"

func main() {

	firstName := []string{
		"Tim",
		"Jenny",
		"Medhi",
		"Marcus",
		"Julian",
		"Sushant",
		"Jose",
		"Harleen",
		"Terminator",
	}

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
		"Howdy",
		"I'll be back.",
	}

	for i, currentEntry := range firstName {
		fmt.Println(currentEntry + ", " + greeting[i])
	}

	fmt.Println()
	fmt.Println("Or this way...")
	fmt.Println()

	for j := 0; j < len(firstName); j++ {
		fmt.Println(firstName[j] + ", " + greeting[j])
	}

}

/*
notes
https://docs.google.com/document/d/17AgfOseB1Pm_9lUxweCA47X-Gemaix3yx1IvoQMYYpI/edit#heading=h.658d4v7m9aao
*/
