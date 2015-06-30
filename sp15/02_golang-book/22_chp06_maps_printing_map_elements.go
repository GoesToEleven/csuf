package main

import "fmt"

/*
The length of a map (found by doing len(x)) can change as we add new items to it.
*/

func main() {

	mapElements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
	}

	mapElements["Li"] = "Lithium"
	mapElements["Be"] = "Beryllium"
	mapElements["B"] = "Boron"
	mapElements["C"] = "Carbon"

	// looking up an element that doesn't exist returns zero value (empty string):
	fmt.Println("Doesn't exist:", mapElements["F"])

	// conditional check of map for element existence
	if value, existence := mapElements["B"]; existence {
		fmt.Println("Exists:", value, existence)
	} else {
		fmt.Println("Doesn't exist:", value, existence)
	}

	// looping through a map (range)
	for key, value := range mapElements {
		fmt.Println("Key:", key, "Value:", value)
	}

}
