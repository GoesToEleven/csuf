package main

import "fmt"

/*
Like arrays and slices, maps can be accessed using brackets.
*/

func main() {
	// ARRAY
	arrayElements := [2]string{
		"Hydrogen",
		"Helium",
	}
	fmt.Println(arrayElements[1])

	// SLICE
	sliceElements := []string{
		"Hydrogen",
		"Helium",
	}
	fmt.Println(sliceElements[1])

	// MAP
	mapElements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
	}
	fmt.Println(mapElements["He"])
}
