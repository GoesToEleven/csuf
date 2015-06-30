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
	go numbers()
	runes()
	fmt.Println("Why doesn't the numbers func print?")
}

// adding "go" puts code onto it's own thread
/*












The answer is down below ...













b/c the program exited before the separate thread could get up and running
*/
