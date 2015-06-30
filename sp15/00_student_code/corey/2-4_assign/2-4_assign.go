package main

import "fmt"

type Name struct {
	first string
	last  string
}

func main() {
	var message string = "Hello, my name is"
	var greeting *string = &message
	var me = Name{"Corey", "Dihel"}

	fmt.Println(message, me.first, me.last)

	*greeting = "You may call me"

	fmt.Println(message, me.first)
}
