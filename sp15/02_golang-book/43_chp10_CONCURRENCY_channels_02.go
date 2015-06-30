package main

import "fmt"

/*
Channels provide a way for two goroutines to communicate with one another
and synchronize their execution.

Let's add another sender to the program and see what happens.
*/

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// this function is new
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		// time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

/*
NICE TO KNOW
Channel Direction

We can specify a direction on a channel type thus restricting it
to either sending or receiving. For example pinger's function signature
can be changed to this:

func pinger(c chan<- string)

Now c can only be sent to. Attempting to receive from c will result in
a compiler error. Similarly we can change printer to this:

func printer(c <-chan string)

A channel that doesn't have these restrictions is known as bi-directional.
A bi-directional channel can be passed to a function that takes send-only
or receive-only channels, but the reverse is not true.
*/
