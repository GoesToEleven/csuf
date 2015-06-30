package main

import "fmt"

// FOR
// -- clause
// ---- init; cond; post
// -- condition
// -- range

func main() {

	i := 0
	for {
		fmt.Println(i)
		i++
	}

	// infinite loops can be used as listeners for events
	// a lot of UI's use them to listen for, and then respond to, user events
	// these are also known as message pumps

	/*

	  for {
	    if someBool {
	      do this
	    } else {
	      do something else
	    }
	  }

	*/
}
