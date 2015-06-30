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
	switch "Jim" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("Have you no friends?")
	}
}
