package main

import "errors"
import "fmt"

func main() {
	err := errors.New("error message")
	fmt.Println(err)
}
