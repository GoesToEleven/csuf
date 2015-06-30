package main

import "fmt"


var done = make(chan bool)

func numbers() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	done <- true
}

func runes() {
	for i := 97; i <= 122; i++ {
		fmt.Println(string(i))
	}

}

func main() {

	go numbers()
	runes()
	fmt.Println("Now the numbers func prints")
	<-done
}

// adding "go" puts code onto it's own thread
