package main

import "fmt"

func main() {

	mySlice := []string{"Rikhil", "Akashdeep"}
	myOtherSlice := []string{"Ishan", "Jenny", "Terminator", "Medhi", "Marcus"}

	mySlice = append(mySlice, myOtherSlice...)

	for _, currentEntry := range mySlice {
		fmt.Println(currentEntry)
	}

	mySlice = append(mySlice[:4], mySlice[5:]...)

	fmt.Println("")

	for _, currentEntry := range mySlice {
		fmt.Println(currentEntry)
	}

}

/*
notes
https://docs.google.com/document/d/17AgfOseB1Pm_9lUxweCA47X-Gemaix3yx1IvoQMYYpI/edit#heading=h.658d4v7m9aao
*/
