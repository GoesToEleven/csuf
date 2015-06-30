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
	case bool:
		return "bool"
	default:
		return "other"
	}
}

func main() {
	fmt.Println("true && true", true && true)
	fmt.Println("true && false", true && false)
	fmt.Println("true || true", true || true)
	fmt.Println("true || false", true || false)
	fmt.Println("!true", !true)

	fmt.Println("********")

	fmt.Println(switchOnType(true && true))
	fmt.Println(switchOnType(true && false))
	fmt.Println(switchOnType(true || true))
	fmt.Println(switchOnType(true || false))
	fmt.Println(switchOnType(!true))
}

/*
Go is a statically typed programming language.
Variables always have a specific type and that type cannot change.

true && true	  true
true && false	  false
false && true	  false
false && false	false

true || true	  true
true || false	  true
false || true	  true
false || false	false

!true	          false
!false	        true




https://golang.org/ref/spec#Numeric_types
*/
