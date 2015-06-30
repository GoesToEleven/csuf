package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("************")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	fmt.Println("************")

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			// divisible by 2
			fmt.Println(i, "divisible by 2")
		} else if i%3 == 0 {
			// divisible by 3
			fmt.Println(i, "divisible by 3")
		} else if i%4 == 0 {
			// divisible by 4
			// QUESTION FOR YOU: WILL THIS LINE BELOW EVER RUN?
			fmt.Println(i, "divisible by 4")
		} else {
			fmt.Println(i, "this is the else")
		}
	}

}
