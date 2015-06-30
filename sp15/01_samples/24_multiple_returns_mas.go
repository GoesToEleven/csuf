package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func Greet(person Contact) {
	fmt.Println(CreateMessage(person.greeting, person.name))
}

func CreateMessage(greeting, name string) (string, string, int) {
	myNum := 120
	return greeting + " " + name, "\nHey, " + name + "\n", myNum
}

func main() {

	var t = Contact{"Good to see you,", "Medhi"}
	Greet(t)

	u := Contact{"Glad you're in class,", "Sushant"}
	Greet(u)

	v := Contact{}
	v.greeting = "We're learning great things,"
	v.name = "Max"
	Greet(v)
}

// put parentheses around return types and comma separate them
// use comma to separate the multiple returns
