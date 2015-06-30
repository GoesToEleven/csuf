package main

import "fmt"

/*
Goroutines are lightweight.
We can easily create thousands of them.
Let's modify our program to run 1,000,000 goroutines:
*/

func f(n int) {
	for i := 0; i <= 2; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	for i := 0; i < 1000000; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

/*
You may have noticed that when you run this program it seems to run
the goroutines in order rather than simultaneously. Let's change this
in the next program.
*/
