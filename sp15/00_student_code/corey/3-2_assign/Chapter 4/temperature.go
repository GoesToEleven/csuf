package main

import "fmt"

func convertF2C(F float32) float32 {
	return (F - 32) * 5 / 9
}

func main() {
	var F float32
	fmt.Println("Please enter a temperature in Fahrenheit:")
	fmt.Scanf("%f", &F)
	C := convertF2C(F)
	fmt.Println("This temperature in Celsius is:", C)
}
