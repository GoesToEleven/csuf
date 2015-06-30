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

type Contact struct {
	greeting string
	name     string
}

// we'll learn more about interfaces later
func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")

	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("McLeod")

	var t = Contact{"Good to see you,", "Tim"}
	SwitchOnType(t)

	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}

// THIS DOES NOT WORK
// func main() {
// 	switch 7 {
// 	case int:
// 		fmt.Println("int")
// 	case string:
// 		fmt.Println("string")
// 	default:
// 		fmt.Println("unknown")
// 	}
// }
