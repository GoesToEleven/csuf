package main

import "fmt"

type contact struct {
	name     string
	greeting string
}

func main() {

	// troubleshooting
	// What's wrong with this code?
	mySlice := []contact{
		"Tim", "Good morning!",
		"Jenny", "Bonjour!",
		"Medhi", "Buenos dias!",
		"Marcus", "Bongiorno!",
		"Julian", "Ohayo!",
		"Sushant", "Selamat pagi!",
		"Jose", "Gutten morgen!",
	}

	for _, currentEntry := range mySlice {
		fmt.Println(currentEntry.name + ", " + currentEntry.greeting)
	}

}

// ANSWER IS DOWN BELOW
/*
































{"Tim", "Good morning!"},
{"Jenny", "Bonjour!"},
{"Medhi", "Buenos dias!"},
{"Marcus", "Bongiorno!"},
{"Julian", "Ohayo!"},
{"Sushant", "Selamat pagi!"},
{"Jose", "Gutten morgen!"},
*/
