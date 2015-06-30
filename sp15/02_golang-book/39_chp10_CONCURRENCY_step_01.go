package main

import "fmt"

/*
Large programs are often made up of many smaller sub-programs.
For example a web server handles requests made from web browsers
and serves up HTML web pages in response. Each request is handled
like a small program.

It would be ideal for programs like these to be able to run their
smaller components at the same time (in the case of the web server
to handle multiple requests). Making progress on more than one task
simultaneously is known as concurrency.

Go has rich support for concurrency using goroutines and channels.

A goroutine is a function that is capable of running concurrently
with other functions. To create a goroutine we use the keyword go
followed by a function invocation.
*/

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)
}

/*
This program consists of two goroutines. The first goroutine is implicit
and is the main function itself. The second goroutine is created when we
call go f(0).

Normally when we invoke a function our program will execute
all the statements in a function and then return to the next line following
the invocation.

With a goroutine we return immediately to the next line and
don't wait for the function to complete.

This is why the call to the Scanln function has been included;
without it the program would exit before being given the opportunity
to print all the numbers.
*/
