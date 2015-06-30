/*
Add a new method to the Shape interface called perimeter which calculates the perimeter of a shape.
Implement the method for Circle and Rectangle.s
 */

package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Circle struct {
x, y, r float64
}

type Shape interface {
	area() float64
	perimeter() float64
}
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2*l + 2*w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

func (c *Circle) perimeter() float64 {
	return math.Pi * c.r * 2
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.perimeter()
	}
	return area
}

func main() {

	c1 := Circle{0,0,4}
	r1 := Rectangle{0,0,7,7}

	fmt.Println(c1.area(), c1.perimeter())
	fmt.Println(r1.area(), r1.perimeter())

	fmt.Println(totalArea(&c1, &r1), totalPerimeter(&c1, &r1))
}
