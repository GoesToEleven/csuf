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

	// fmt.Println()
	// fmt.Println("Or this way...")
	// fmt.Println()

	// fmt.Println("ACCESS JUST THE VALUE:")
	// val := firstName[0]
	// fmt.Println("val: ", val)
	// fmt.Println()

	// fmt.Println("ACCESS THE VALUE & AN EXISTENCE BOOL:")
	// val, exists := firstName[0]
	// fmt.Println("val: ", val)
	// fmt.Println("exists: ", exists)
	// fmt.Println()

	// fmt.Println("CHECK FOR EXISTENCE, THEN ITERATE:")

	// for i := 0; i < len(firstName); i++ {
	//
	//   if val, exists := myGreeting["Terminator"]; exists {
	//
	// 	fmt.Println(firstName[i] + ", " + greeting[i])
	// }

	// Maps - Shorthand Notation
	// myGreeting ,= map[string]string{
	//   "Tim":        "Good morning!",
	//   "Jenny":      "Bonjour!",
	//   "Medhi":      "Buenos dias!",
	//   "Marcus":     "Bongiorno!",
	//   "Julian":     "Ohayo!",
	//   "Sushant":    "Selamat pagi!",
	//   "Jose":       "Gutten morgen!",
	//   "Harleen":    "Howdy",
	//   "Terminator": "I'll be back.",
	// }

}

/*
notes
https://docs.google.com/document/d/17AgfOseB1Pm_9lUxweCA47X-Gemaix3yx1IvoQMYYpI/edit#heading=h.658d4v7m9aao
*/
