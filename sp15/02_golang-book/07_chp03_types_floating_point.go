package main

import "fmt"

func switchOnType(x interface{}) string {
	switch x.(type) {
	case float32:
		return "float32"
	case float64:
		return "float64"
	case complex64:
		return "complex64"
	case complex128:
		return "complex128"
	case int:
		return "int"
	case uint:
		return "uint"
	// case uint8:
	// 	return "uint8"
	case byte:
		return "byte"
	case uint16:
		return "uint16"
	case uint32:
		return "uint32"
	case uint64:
		return "uint64"
	case int8:
		return "int8"
	case int16:
		return "int16"
	// case int32:
	// 	return "int32"
	case rune:
		return "rune"
	case int64:
		return "int64"
	case string:
		return "string"
	default:
		return "other"
	}
}

func addInt() {
	fmt.Println("1 + 1 =", 1+1, switchOnType(1+1))
}

func addFloat() {
	fmt.Println("1.0 + 1.0 =", 1.0+1.0, switchOnType(1.0+1.0))
}

func addFloatMas() {
	fmt.Println("1 + 1.0 =", 1+1.0, switchOnType(1+1.0))
}

func divideInt() {
	fmt.Println("10 / 3 =", 10/3, switchOnType(10/3))
}

func divideFloat() {
	fmt.Println("10 / 3.0 =", 10/3.0, switchOnType(10/3.0))
}

func divideFloatMas() {
	fmt.Println("10.0 / 3 =", 10.0/3, switchOnType(10.0/3))
}

func main() {
	addInt()
	addFloat()
	addFloatMas()
	divideInt()
	divideFloat()
	divideFloatMas()
}

/*
Go is a statically typed programming language.
Variables always have a specific type and that type cannot change.

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

In addition to numbers there are several other values which can be represented: “not a number” (NaN, for things like 0/0) and positive and negative infinity. (+∞ and −∞)

Go has two floating point types: float32 and float64
(also often referred to as single precision and double precision respectively)

Generally we should stick with float64 when working with floating point numbers.

https://golang.org/ref/spec#Numeric_types

*/
