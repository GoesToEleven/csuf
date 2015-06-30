package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main() {
	someOtherName := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(someOtherName))
}

/*
A function is an independent section of code that maps zero or more input parameters
to zero or more output parameters. Functions (also known as procedures or subroutines)
are often represented as a black box: (the black box represents the function)

SIGNATURE
the parameters we put the return type. Collectively the parameters and the return type
are known as the function's signature

RETURN
return statement causes the function to immediately stop and return the value
after it to the function that called this one

NAME RETURN TYPE
We can also name the return type:

func f2() (r int) {
    r = 1
    return
}


SCOPE
Functions don't have access to anything in the calling function. This won't work:

func f() {
    fmt.Println(x)
}

func main() {
    x := 5
    f()
}

WE NEED TO DO EITHER THIS:

func f(x int) {
    fmt.Println(x)
}

func main() {
    x := 5
    f(x)
}

OR THIS:

var x int = 5

func f() {
    fmt.Println(x)
}

func main() {
    f()
}

*/
