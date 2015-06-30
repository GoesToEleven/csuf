package main

import "fmt"

/*
* no default fall through
* fall through is optional
* -- you can specify fall through by explicitly stating it
* expression not needed
* -- if no expression provided, go checks for the first case that evals to true
* -- makes the switch operate like if/if else/else
* cases can be expressions
* you can switch on types
* -- normally we switch on value of variable
* -- go allows you to switch on type of variable
* ---- if it's an int you can do one thing, if it's a string you can do another
 */

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")
	}
}
