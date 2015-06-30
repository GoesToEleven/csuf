package main

import (
	"fmt"
	"time"
)

/*
The select statement is often used to implement a timeout
*/

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case <-time.After(time.Second * 1):
				fmt.Println("timeout")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

/*
time.After creates a channel and after the given duration will send
the current time on it. (we weren't interested in the time so we didn't
store it in a variable)
*/
