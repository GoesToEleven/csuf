package main

import "fmt"

func mySwitch(x interface{}) {
	switch x {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}

func main() {
	mySwitch(3)
	mySwitch(2)
	mySwitch(5)
}
