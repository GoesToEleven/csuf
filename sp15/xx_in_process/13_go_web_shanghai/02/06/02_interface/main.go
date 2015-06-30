package main

import "fmt"

func main() {
	// define a as empty interface
	var a interface{}
	var i int = 5
	s := "Hello world"
	// a can store value of any type
	a = i
	fmt.Println(a)
	a = s
	fmt.Println(a)
}
