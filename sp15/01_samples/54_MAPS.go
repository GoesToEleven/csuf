package main

import "fmt"

/*

basic info
--key:value
--maps keys to values
--called “dictionaries” in some languages
--built into the language (not an additional library you must import) so they’re first-class citizens
Map Keys
--need to be unique
--the type used for a key needs to have the equality operator defined for it
----the type must allow equals comparisons
----can’t use these types
------slice
------map
Maps are Reference Types
--maps are reference types
--they behave like pointers
----when you pass a map variable to a function
------any changes to that mapped variable in the function
------change that original mapped variable outside the function
Maps are Not Thread Safe
--best to avoid using maps concurrently

*/

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

	fmt.Println(myGreeting["Jenny"])
}
