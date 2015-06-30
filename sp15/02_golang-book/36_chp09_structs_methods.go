package main

import (
	"fmt"
	"math"
)

// Circle is a struct data type
type Circle struct {
	x, y, r float64
}

// WITHOUT method
/*
func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}
*/

// WITH method
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{0, 0, 5}
	/* In between the keyword func and the name of the function
	   we've added a “receiver”. The receiver is like a parameter –
	   it has a name and a type – but by creating the function in this way
	   it allows us to call the function using the . operator:
	*/
	fmt.Println(c.area())
}

/*
we no longer need the & operator (Go automatically knows to pass
a pointer to the circle for this method) and because this function
can only be used with Circles we can rename the function to just area
*/
