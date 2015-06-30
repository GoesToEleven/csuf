package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func main() {

	mySlice := []contact{
		{"Good to see you", "Medhi"},
		{"Glad you're in class", "Sushant"},
	}

	for _, currentEntry := range mySlice {
		fmt.Println(currentEntry.greeting + ", " + currentEntry.name)
	}

	/*
		  THIS IS ALL FOR REFERENCE - WHAT WE HAD BEFORE:

			     var t = Contact{"Good to see you,", "Medhi"}

			     u := Contact{"Glad you're in class,", "Sushant"}

			     v := Contact{}
			     v.greeting = "We're learning great things,"
			     v.name = "Marcus"


			  func Greet(person Contact) {
			  	fmt.Println(CreateMessage(person.greeting, person.name))
			  }

			  func CreateMessage(greeting, name string) string {
			  	return greeting + " " + name
			  }

	*/
}
