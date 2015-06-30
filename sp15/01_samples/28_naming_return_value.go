package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func Greet(person Contact) {
	myGreetingMas, myNameMas := CreateMessage(person.greeting, person.name)
	fmt.Print(myGreetingMas)
	fmt.Print(myNameMas)
}

// YOU CAN ALSO NAME THE RETURN VALUES
func CreateMessage(greeting, name string) (myGreeting string, myName string) {
	myGreeting = greeting + " " + name
	myName = "\nHey, " + name + "\n"
	return
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
