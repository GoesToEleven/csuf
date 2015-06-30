package main

import "fmt"

// TYPES
type myPhrase string

// FUNCS
func rephraseUno(m myPhrase) {
	m = "Uno - Bali in the summer."
}

func rephraseDos(m *myPhrase) {
	*m = "Dos - Bali in the summer."
}

func (m *myPhrase) rephraseTres() {
	*m = "Tres - Uluwatu all afternoon."
}

// func (m myPhrase) rephraseQuatro(s string) {
// 	m = s
// }

func main() {
	// VARIABLES
	var toddSays myPhrase
	// TYPES CAN CONTAIN DATA
	toddSays = "Cool waves and warm breezes."
	fmt.Println(toddSays)
	// METHOD pass by value
	rephraseUno(toddSays)
	fmt.Println(toddSays)
	// METHOD pass by reference
	rephraseDos(&toddSays)
	fmt.Println(toddSays)
	// METHOD attached to a type
	toddSays.rephraseTres()
	fmt.Println(toddSays)
	// METHOD attached to a type
	// toddSays.rephraseQuatro("Quatro - Uluwatu all afternoon.")
	// fmt.Println(toddSays)

}
