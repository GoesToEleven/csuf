// STEP 1 - GOAL:
// to be able to add *any* string to the end of what is printed

package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

// create your own printer type
type MyPrinterType func(string)

// pass in a function
// using your printer type
func Greet(person Contact, myWassa MyPrinterType) {
	myGreetingMas, myNameMas := CreateMessage(person.name, person.greeting, "howdy")
	myWassa(myGreetingMas)
	myWassa(myNameMas)
}

func CreateMessage(name string, greeting ...string) (myGreeting string, myName string) {
	myGreeting = greeting[1] + " " + name
	myName = "\nHey, " + name + "\n"
	return
}

// create some functions you might want to pass in
func myPrint(s string) {
	fmt.Print(s)
}

func myPrintln(s string) {
	fmt.Println(s)
}

// STEP 2:
// we could create this function, HOWEVER ...
// we would then need to modify our custom type (MyPrinterType) to take two params
// so this func below IS NOT a good solution
func myPrintCustom(s string, custom string) {
	fmt.Println(s + custom)
}

// STEP 3:
// a better solution is to create a func that returns a func
// this is also known as "a function factory"
func myPrintFunction(custom string) MyPrinterType {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

// STEP 4:
// use your function
// pass in a function as an argument to another function
func main() {

	var t = Contact{"Good to see you,", "Tim"}
	// had this before:
	// Greet(t, myPrint)
	Greet(t, myPrintFunction("!!!"))

	u := Contact{"Glad you're in class,", "Jenny"}
	// had this before:
	// Greet(u, myPrint)
	Greet(u, myPrintFunction("???"))

	v := Contact{}
	v.greeting = "We're learning great things,"
	v.name = "Julian"
	// had this before:
	// Greet(v, myPrintln)
	Greet(v, myPrintFunction("^^^"))
}
