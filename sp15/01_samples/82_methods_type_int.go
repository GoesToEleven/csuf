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
	fmt.Println("Rectangle is: ", oneRectangle)

	// TYPES CAN USE METHODS
	fmt.Println("Rectangle area is: ", oneRectangle.area())
}
