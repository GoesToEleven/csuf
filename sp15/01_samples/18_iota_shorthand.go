package main

import "fmt"

const (
	PI       = 3.14
	Language = "Go"
)

const (
	A = iota // 0
	B        // 1
	C        // 2
)

func main() {
	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

// iota represents successive untyped integer constants
