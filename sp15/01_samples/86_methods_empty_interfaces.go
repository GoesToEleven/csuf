package main

import "fmt"

// an EMPTY INTERFACE allows you to pass any type into a function/method

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")

	}
}

// the EMPTY INTERFACE is being declared when SwitchOnType is being declared
// question: is the empty interface a parameter, an argument, both, or neither?

func main() {
	SwitchOnType(7)
	SwitchOnType("McLeod")
	SwitchOnType(7.4)
}
