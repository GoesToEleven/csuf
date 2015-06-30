package main

import "fmt"

func main() {

	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")

}

/*
A panic generally indicates a programmer error (for example attempting to access
an index of an array that's out of bounds, forgetting to initialize a map, etc.)
or an exceptional condition that there's no easy way to recover from.
(Hence the name “panic”)

We can handle a run-time panic with the built-in recover function.
recover stops the panic and returns the value that was passed to the call to panic.
We might be tempted to use it like this:

package main

import "fmt"

func main() {
   panic("PANIC")
   str := recover()
   fmt.Println(str)
}

But the call to recover will never happen in this case because the call to panic
immediately stops execution of the function. Instead, do it like the code which
isn't commented out above.

*/
