package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func Greet(contact Contact) {
	myGreetingMas, myNameMas := CreateMessage(contact.name, contact.greeting, "howdy")
	fmt.Print(myGreetingMas)
	fmt.Print(myNameMas)
}

// VARIADIC FUNCTIONS
// a variable number of parameters of a certain type
// useful when we don't know how many parameters are going to be passed in
func CreateMessage(name string, greeting ...string) (myGreeting string, myName string) {
	fmt.Println(len(greeting))
	myGreeting = greeting[1] + " " + name
	myName = "\nHey, " + name + "\n"
	return
}

func main() {

	var t = Contact{"Good to see you,", "Tim"}
	Greet(t)

	u := Contact{"Glad you're in class,", "Jenny"}
	Greet(u)

	v := Contact{}
	v.greeting = "We're learning great things,"
	v.name = "Julian"
	Greet(v)
}
