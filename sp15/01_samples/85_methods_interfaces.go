package main

import "fmt"

type contact struct {
	name     string
	greeting string
}

type contacts []contact

// (1) - declare a type
// (2) - say that it's an interface
// (3) - declare the method(s)
type renamable interface {
	rename(newName string)
}

func (c *contact) rename(newName string) {
	c.name = newName
}

// any func that implements the same methods as the interface
// will be implementing that interface
func renamer(r renamable) {
	r.rename("Bono")
}

func main() {

	myContacts := contacts{
		{"Ying", "Hello!"},
		{"Colin", "Buenos Dias!"},
		{"Brandon", "Bonjour!"},
		{"Shaymas", "Top of the morning!"},
		{"Dan", "Good day!"},
	}

	for _, currentEntry := range myContacts {
		fmt.Println(currentEntry.name + " - " + currentEntry.greeting)
	}
	fmt.Println()

	for _, currentEntry := range myContacts {
		// currentEntry.rename("Bono")
		renamer(&currentEntry)
		fmt.Println(currentEntry.name + " - " + currentEntry.greeting)
	}

}
