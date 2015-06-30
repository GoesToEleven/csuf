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

	myFriendsName := "Medhi"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Both of your names start with M")
	case myFriendsName == "Julian":
		fmt.Println("Wassup Julian")
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Sushant")
	}
}

// HAD THIS BEFORE
//   switch "Marcus" {
//   case "Tim":
//     fmt.Println("Wassup Tim")
//   case "Jenny":
//     fmt.Println("Wassup Jenny")
//   case "Marcus":
//     fmt.Println("Wassup Marcus")
//     fallthrough
//   case "Medhi":
//     fmt.Println("Wassup Medhi")
//   case "Julian":
//     fmt.Println("Wassup Julian")
//   case "Sushant":
//     fmt.Println("Wassup Sushant")
//   }
