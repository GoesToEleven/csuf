package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Let's add some delay to the function using time.Sleep and rand.Intn:
*/

func f(n int) {
	for i := 0; i < 5; i++ {
		fmt.Println(n, ":", string(i+65))
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

/*
f prints out the letters from A to J, waiting between 0 and 250 ms
after each one. You should see now how the goroutines run simultaneously.
*/
