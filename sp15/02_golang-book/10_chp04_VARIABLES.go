package main

import "fmt"

// this works
var x string = "Hello World"

// this does not work
// you cannot use shorthand notation ourside of a function
// x := "Hello world"

func main() {
	fmt.Println(x)
}
