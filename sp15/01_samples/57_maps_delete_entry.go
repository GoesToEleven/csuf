package main

import "fmt"

func main() {

	// Maps - Shorthand Notation
	myGreeting := map[string]string{
		"Tim":        "Good morning!",
		"Jenny":      "Bonjour!",
		"Medhi":      "Buenos dias!",
		"Marcus":     "Bongiorno!",
		"Julian":     "Ohayo!",
		"Sushant":    "Selamat pagi!",
		"Jose":       "Gutten morgen!",
		"Harleen":    "Howdy",
		"Terminator": "I'll be back.",
	}

	fmt.Println(myGreeting["Terminator"])

	myGreeting["Terminator"] = "There are 10 types of people ..."

	fmt.Println(myGreeting["Terminator"])
	fmt.Println()

	// terminate the terminator
	// delete(myGreeting, "Terminator")

	fmt.Println("ACCESS JUST THE VALUE:")
	val := myGreeting["Terminator"]
	fmt.Println("val: ", val)
	fmt.Println()

	fmt.Println("ACCESS THE VALUE & AN EXISTENCE BOOL:")
	val, exists := myGreeting["Terminator"]
	fmt.Println("val: ", val)
	fmt.Println("exists: ", exists)
	fmt.Println()

	fmt.Println("CHECK FOR EXISTENCE:")
	if val, exists := myGreeting["Terminator"]; exists {
		fmt.Println("You can't kill the governor")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("I have been terminated.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

}

/* documentation:
https://docs.google.com/document/d/17AgfOseB1Pm_9lUxweCA47X-Gemaix3yx1IvoQMYYpI/edit#heading=h.hpcmyxdwosk4
*/
