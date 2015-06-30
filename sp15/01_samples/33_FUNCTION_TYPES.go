package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

// FUNCTIONS ARE TYPES
// functions in go are types
// functions behave as types in go
// you can pass functions around just as you'd pass types around
// pass functions just like any other argument / parameter

// STEP 1:
// create a function that takes as a parameter a function
func Greet(person Contact, myWassa func(string)) {
	myGreetingMas, myNameMas := CreateMessage(person.name, person.greeting, "howdy")
	// had this before:
	// fmt.Print(myGreetingMas)
	// fmt.Print(myNameMas)
	myWassa(myGreetingMas)
	myWassa(myNameMas)
}

func CreateMessage(name string, greeting ...string) (myGreeting string, myName string) {
	myGreeting = greeting[1] + " " + name
	myName = "\nHey, " + name + "\n"
	return
}

// STEP 2:
// create some functions you might want to pass in
func myPrint(s string) {
	fmt.Print(s)
}

// STEP 2:
// create some functions you might want to pass in
func myPrintln(s string) {
	fmt.Println(s)
}

// STEP 3:
// pass in functions (passing a function as an argument to another function)
func main() {

	var t = Contact{"Good to see you,", "Tim"}
	Greet(t, myPrint)

	u := Contact{"Glad you're in class,", "Jenny"}
	Greet(u, myPrint)

	v := Contact{}
	v.greeting = "We're learning great things,"
	v.name = "Julian"
	Greet(v, myPrintln)
}

// see this link for a nice diagram of the above process:
// https://docs.google.com/document/d/17AgfOseB1Pm_9lUxweCA47X-Gemaix3yx1IvoQMYYpI/edit#bookmark=id.8072tlk5e19t
