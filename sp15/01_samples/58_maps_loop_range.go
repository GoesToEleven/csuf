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

	for key, value := range myGreeting {
		fmt.Println("Key:", key, "Value:", value)
	}

	// run this multiple times and watch the results shift!

}

/* documentation:
https://docs.google.com/document/d/17AgfOseB1Pm_9lUxweCA47X-Gemaix3yx1IvoQMYYpI/edit#heading=h.hpcmyxdwosk4
*/
