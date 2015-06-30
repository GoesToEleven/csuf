package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
