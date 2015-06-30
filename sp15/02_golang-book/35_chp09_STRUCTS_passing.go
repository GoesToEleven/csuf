package main

import "fmt"

// Circle is a struct data type
type Circle struct {
	x, y, r float64
}

func passByVal(c Circle) {
	fmt.Println("Inside passByVal", c.r)
	c.r = 10
	fmt.Println("Inside passByVal", c.r)
}

func passByRef(c *Circle) {
	fmt.Println("Inside passByRef", c.r)
	c.r = 10
	fmt.Println("Inside passByRef", c.r)
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println("beginning: ", c.x, c.y, c.r)
	passByVal(c)
	fmt.Println("after passByVal: ", c.x, c.y, c.r)
	passByRef(&c)
	fmt.Println("after passByRef", c.x, c.y, c.r)
	c.x = 10
	c.y = 5
	fmt.Println("after dot notation: ", c.x, c.y, c.r)
	fmt.Println("***********")
	fmt.Println(&c)
	addLoc := &c.x
	fmt.Println(addLoc)

}
