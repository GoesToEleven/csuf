package main

import (
	"fmt"
)

type my_operator func(int, int) int

func my_operation(num1 int, num2 int, op my_operator) int {
	result := op(num1, num2)
	if result < 0 {
		fmt.Println("The result is negative")
	} else {
		fmt.Println("The result is positive")
	}
	return result
}

func operator_creation() (my_operator, my_operator, my_operator, my_operator) {
	return func(a int, b int) int {
			return a + b
		}, func(a int, b int) int {
			return a - b
		}, func(a int, b int) int {
			return a * b
		}, func(a int, b int) int {
			return a / b
		}
}

func main() {
	add, subtract, multiply, divide := operator_creation()
	a, b := 6, 3
	var result int
	operation := "multiply"
	if a < b {
		operation = "subtract"
	}
	switch operation {
	case "add":
		result = my_operation(a, b, add)
	case "subtract":
		result = my_operation(a, b, subtract)
	case "multiply":
		result = my_operation(a, b, multiply)
	case "divide":
		result = my_operation(a, b, divide)
	}
	fmt.Println("The result is:", result)
}
