package main

import "fmt"

const (
	PI       = 3.14
	Language = "Go"
	A        = iota // 2
	B        = iota // 3
	C        = iota // 4
)

func main() {
	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

// iota represents successive untyped integer constants
