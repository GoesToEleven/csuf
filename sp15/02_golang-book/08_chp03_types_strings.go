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

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println(switchOnType(len("Hello World")))
	fmt.Println("**********************")

	fmt.Println("Hello World"[1])
	fmt.Println(string("Hello World"[1]))
	fmt.Println(switchOnType("Hello World"[1]))
	fmt.Println("**********************")

	fmt.Println("Hello " + "World")
	fmt.Println(switchOnType("Hello " + "World"))
	fmt.Println("**********************")

	fmt.Println("Hello World"[6:])
	fmt.Println("Hello World"[:5])
	fmt.Println("Hello World"[3:8])
	fmt.Println("**********************")
}

/*
Go is a statically typed programming language.
Variables always have a specific type and that type cannot change.

Go strings are made up of individual bytes, usually one for each character.
Characters from other languages like Chinese are represented by more than one byte



https://golang.org/ref/spec#Numeric_types
*/
