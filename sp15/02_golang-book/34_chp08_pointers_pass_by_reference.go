package main

import "fmt"

func zero(xPtr *int) {
	fmt.Println("xPtr: ", xPtr)
	fmt.Println("*xPtr: ", *xPtr)
	fmt.Println("***********")
	*xPtr = 0
}

func main() {
	x := 5
	fmt.Println("x: ", x)
	fmt.Println("&x: ", &x)
	fmt.Println("***********")
	zero(&x)
	fmt.Println(x) // x is 0
}

/*
Pointers reference a location in memory where a value is stored rather than
the value itself. (They point to something else) By using a pointer (*int)
the zero function is able to modify the original variable.

In Go a pointer is represented using the * (asterisk) character followed by
the type of the stored value. In the zero function xPtr is a pointer to an int.

* is also used to “dereference” pointer variables. Dereferencing a pointer gives
us access to the value the pointer points to. When we write *xPtr = 0 we are saying
“store the int 0 in the memory location xPtr refers to”. If we try xPtr = 0 instead
we will get a compiler error because xPtr is not an int it's a *int,
which can only be given another *int.

Finally we use the & operator to find the address of a variable.
&x returns a *int (pointer to an int) because x is an int.
This is what allows us to modify the original variable.
&x in main and xPtr in zero refer to the same memory location.

*/
