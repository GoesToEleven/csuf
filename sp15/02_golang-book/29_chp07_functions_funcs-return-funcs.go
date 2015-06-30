package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
}

/*
One way to use closure is by writing a function which returns another function
which – when called – can generate a sequence of numbers.

makeEvenGenerator returns a function which generates even numbers.
Each time it's called it adds 2 to the local i variable which
– unlike normal local variables – persists between calls.
*/
