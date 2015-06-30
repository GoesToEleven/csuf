package main

import "fmt"

// TYPE
type rectangle struct {
	length, width int
}

// TYPES CAN CONTAIN METHODS & RETURN
func (r rectangle) area() int {
	return r.length * r.width
}

func main() {

	// TYPES CAN CONTAIN DATA
	oneRectangle := rectangle{5, 4}
	anotherRectangle := rectangle{7, 3}
	otroMas := rectangle{9, 5}
	fmt.Println("Rectangle is: ", oneRectangle)
	fmt.Println("Rectangle is: ", anotherRectangle)
	fmt.Println("Rectangle is: ", otroMas)

	// TYPES CAN USE METHODS
	fmt.Println("Rectangle area is: ", oneRectangle.area())
	fmt.Println("Rectangle area is: ", anotherRectangle.area())
	fmt.Println("Rectangle area is: ", otroMas.area())
}
