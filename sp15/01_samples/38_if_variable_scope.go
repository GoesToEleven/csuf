package main

import "fmt"

func myConditional(myCheck bool) {

	// we could do this, HOWEVER
	// it would have a wider scope
	// myFavName := "Oliver"

	// instead, we can do this
	// and narrow the scope
	if myFavName := "Oliver"; myCheck {
		fmt.Println("This ran, " + myFavName)
	}

	// NOTE: we can't have this b/c outside scope
	// fmt.Println("Wassup, " + myFavName)

	fmt.Println("Wassup.")

}

func main() {
	myConditional(true)
}
