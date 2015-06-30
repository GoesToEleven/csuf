package main

import (
	"fmt"
	"time"
)

func main() {
	t := make(chan time.Time)

	go func() {
		for {
			mytime := <-t
			fmt.Println(mytime.Format("3:04:05pm"))
		}
	}()

	go func() {
		for {
			t <- <-time.After(time.Second * 1)
		}
	}()

	var input string
	fmt.Scanln(&input)
}
