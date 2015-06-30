package main

import "fmt"

func numbers() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func runes() {
	for i := 97; i <= 122; i++ {
		fmt.Println(string(i))
	}

}

func main() {
	numbers()
	runes()
}

// this example has no concurency
// control flow is sequential, top to bottom
