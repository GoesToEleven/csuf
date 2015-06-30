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

	fmt.Println(len(mapElements))

	mapElements["Li"] = "Lithium"
	mapElements["Be"] = "Beryllium"
	mapElements["B"] = "Boron"
	mapElements["C"] = "Carbon"

	fmt.Println(len(mapElements))

	delete(mapElements, "B")

	fmt.Println(len(mapElements))
}
