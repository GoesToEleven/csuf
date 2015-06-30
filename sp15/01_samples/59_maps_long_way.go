package main

import "fmt"

func main() {

	/*
	   --declaration
	   ----declare a variable of map type
	   ------specify key type
	   ------specify value type
	   make your map
	   --“make”  - special keyword
	   ----makes the map; allocates memory; initializes the map
	   ----puts 0 or empty string into values
	   ----if you don’t use make, then variable values are nil - empty, doesn’t have any items
	*/

	var myGreeting map[string]string
	myGreeting = make(map[string]string)

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(myGreeting["Harleen"])

}
