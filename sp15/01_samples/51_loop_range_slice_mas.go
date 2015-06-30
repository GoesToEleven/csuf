package main

import "fmt"

// FOR
// -- clause
// ---- init; cond; post
// -- condition
// -- range
// ---- works on these types:
// ------ slice or array
// ------ string
// -------- gives us a rune (code point to UTF-8 character)
// ------ map
// -------- key:value
// ------ channel
// -------- a channel is a way to communicate between threads (different go routines)
// -------- you can use the "for range" to read off of a channel continously

func main() {

	mySlice := []string{
		"Good morning",
		"Bonjour",
		"Buenos dias",
		"Bongiorno",
		"Ohayo",
		"Selamat pagi",
		"Gutten morgen",
	}

	for _, currentEntry := range mySlice {
		fmt.Println(currentEntry)
	}

}
