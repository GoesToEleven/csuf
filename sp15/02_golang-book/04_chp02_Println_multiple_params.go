package main

import "fmt"

func main() {
	b, e := fmt.Println("Hello World", "dog")
	fmt.Println("bytes written:", b)
	fmt.Println("error:", e)
}

/*
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.
*/
