package main

import "fmt"

/*
A map is an unordered collection of key-value pairs.
Also known as an associative array, a hash table, or a dictionary,
maps are used to look up a value by its associated key.
*/

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"

	fmt.Println(elements["He"])
}
