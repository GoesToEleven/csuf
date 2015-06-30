package main

import "fmt"

func main() {

	greeting := "I'll be back."

	fmt.Print("[5:7] ")
	fmt.Println(greeting[5:7])
	fmt.Print("[:4] ")
	fmt.Println(greeting[:4])
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:])
	// try others!
	// fmt.Print("[:] ")
	// fmt.Println(greeting[:])
}

/*
notes
https://docs.google.com/document/d/17AgfOseB1Pm_9lUxweCA47X-Gemaix3yx1IvoQMYYpI/edit#heading=h.658d4v7m9aao
*/
