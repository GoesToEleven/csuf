package main

import "fmt"

// Person is a struct which holds name
type Person struct {
	Name string
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// NOT THIS WAY
/*
type Android struct {
    Person Person
    Model string
}
*/

// THIS WAY
// Android is a struct which holds a Person struct and a model
// android "is a" person; not android "has a" person
// you're embedding one type in another type
type Android struct {
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Person.Name = "Shamay"
	a.Person.talk()
	a.Name = "Julian"
	a.talk()
}
