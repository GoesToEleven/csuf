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

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return c.r * c.r * math.Pi
}

func (c *Circle) perimeter() float64 {
	return 2 * c.r * math.Pi
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 5, 5}
	fmt.Println("The area of the circle is:", c.area())
	fmt.Println("The area of the rectangle is:", r.area())
	fmt.Println("The perimeter of the circle is:", c.perimeter())
	fmt.Println("The perimeter of the rectangle is:", r.perimeter())
}
