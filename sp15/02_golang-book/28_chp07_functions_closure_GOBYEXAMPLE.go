package main

import "fmt"

/*
This function intSeq returns another function, which we define anonymously
in the body of intSeq. The returned function closes over the variable i to form a closure.
*/

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

/*
   We call intSeq, assigning the result (a function) to nextInt.
   This function value captures its own i value, which
   will be updated each time we call nextInt.

   See the effect of the closure by calling nextInt a few times.

   To confirm that the state is unique to that particular function, create and test a new one.
*/
