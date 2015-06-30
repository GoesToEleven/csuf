package main

import "fmt"

func main() {
	const Meter float32 = 0.3048
	var F float32
	fmt.Println("Please enter a length in Feet:")
	fmt.Scanf("%f", &F)
	fmt.Println("This length in Meters is:", F*Meter)
}
