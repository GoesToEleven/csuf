package main

import "fmt"

func myConditional(myCheck bool) {

	if myCheck {
		fmt.Println("This ran")
	}

	fmt.Println("Wassup.")

}

func main() {
	myConditional(false)
}
