package main

/*
This is known as a “package declaration”.
Every Go program must start with a package declaration.
Packages are Go's way of organizing and reusing code.
*/

import "fmt"

/*
The import keyword is how we include code from other packages.
The fmt package (shorthand for format) implements formatting for input and output.
What do you think the fmt package's files would contain at the top of them?
*/

func main() {
	fmt.Println("Hello World")
}

/*
Functions are the building blocks of a Go program. They have inputs, outputs
and a series of steps called statements which are executed in order. All
functions start with the keyword func followed by the name of the function
(main in this case), a list of zero or more “parameters” surrounded by
parentheses, an optional return type and a “body” which is surrounded by curly
braces. This function has no parameters, doesn't return anything and has only
one statement. The name main is special because it's the function that gets
called when you execute the program.

fmt.Println("Hello World")
This statement is made of three components. First we access another function
inside of the fmt package called Println (that's the fmt.Println piece, Println
means Print Line). Then we create a new string that contains Hello World and
invoke (also known as call or execute) that function with the string as the
first and only argument.
*/

// this is all from www.golang-book.com
