/*
The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
Write a recursive function which can find fib(n).
 */

package main

import (
	"fmt"
)

func fib(x int) int {
	if x < 0 {
		defer func() {
			fmt.Println(recover())
		}()
		panic("don't do that")
	} else if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return fib(x-1) + fib(x-2)
	}
}

func main() {
	for i := -10; i <= 10; i++{
		fmt.Println(i,fib(i))
	}
}
