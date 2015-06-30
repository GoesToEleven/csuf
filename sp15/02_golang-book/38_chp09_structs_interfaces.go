package main

import "fmt"
import "math"

/*
an interface defines a list of methods that a type must have
in order to “implement” the interface.
*/

// Shape is an interface which requires an area() method that returns float64
type Shape interface {
	area() float64
}

// Circle is a struct
type Circle struct {
	x, y, r float64
}

// Rectangle is a struct
type Rectangle struct {
	l, w float64
}

// Circle implements the Shape interface
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Rectangle implements the Shape interface
func (r *Rectangle) area() float64 {
	return r.l * r.w
}

// we can use interface types as arguments to functions
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{1, 3, 5}
	fmt.Println(c.area())
	r := Rectangle{2, 10}
	fmt.Println(r.area())
	fmt.Println(totalArea(&c, &r))
}
