/*
Write your own Sleep function using time.After.
 */

package main

import (
	"fmt"
	"time"
)

func Sleep( timeInMs time.Duration ){
	 <- time.After(timeInMs * time.Second / 1000 )
}

func main() {
	fmt.Println("Starting Sleep")
	Sleep(1000);
	fmt.Println("Ending Sleep")
}
