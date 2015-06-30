package main

import "fmt"

func main() {

	mySlice := []int{1, 5, 15, 20, 25, 30}
	myOtherSlice := []int{50, 70, 90}

	mySlice = append(mySlice, myOtherSlice...)

	for i, currentEntry := range mySlice {
		fmt.Println(i, " - ", currentEntry)
	}

}

/*
notes
https://docs.google.com/document/d/17AgfOseB1Pm_9lUxweCA47X-Gemaix3yx1IvoQMYYpI/edit#heading=h.658d4v7m9aao
*/
