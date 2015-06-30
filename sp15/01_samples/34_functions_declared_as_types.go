package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

// STEP 1:
// declare a type
type MyPrinterType func(string)

// use the type
// we had this before:
// func Greet(person Contact, myWassa func(string)) {
func Greet(person Contact, myWassa MyPrinterType) {
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

func myPrintln(s string) {
	fmt.Println(s)
}

// STEP 3:
// pass in functions (passing a function as an argument to another function)
func main() {

	var t = Contact{"Good to see you,", "Tim"}
	Greet(t, myPrint)

	u := Contact{"Glad you're in class,", "Jenny"}
	Greet(u, myPrintln)

	v := Contact{}
	v.greeting = "We're learning great things,"
	v.name = "Julian"
	Greet(v, func(s string) {
		fmt.Println(s)
	})
}

// see this link for a nice diagram of the above process:
// https://docs.google.com/document/d/17AgfOseB1Pm_9lUxweCA47X-Gemaix3yx1IvoQMYYpI/edit#bookmark=id.8072tlk5e19t
