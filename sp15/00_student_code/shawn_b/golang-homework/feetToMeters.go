package main

import(
	"fmt"
)

func main() {
	fmt.Print("Enter a length in feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 0.3048

	fmt.Println("Meters =",output)
}
