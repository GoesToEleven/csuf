package main

import "fmt"

func main() {

	mySlice := []int{1, 5, 15, 20, 25, 30}

	for i, currentEntry := range mySlice {
		fmt.Println(i, " - ", currentEntry)
	}

}

/*
notes
https://docs.google.com/document/d/17AgfOseB1Pm_9lUxweCA47X-Gemaix3yx1IvoQMYYpI/edit#heading=h.658d4v7m9aao
*/
