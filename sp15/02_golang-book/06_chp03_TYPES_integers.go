package main

import "fmt"

func switchOnType(x interface{}) string {
	switch x.(type) {
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
		return "unknown"
	}
}

func main() {

	// run either casting or noCasting, but not both

	// NO CASTING
	// myint := -255
	// myuint := 255
	// myuint8 := 255
	// mybyte := 255
	// myuint16 := 65535
	// myuint32 := 4294967295
	// /* myuint64 := 18446744073709551615 */
	// myint8 := -255
	// myint16 := -65535
	// myint32 := -4294967295
	// myrune := -4294967295
	// myint64 := -9223372036854775807

	// CASTING
	myUncast := -255
	myint := int(-255)
	myuint := uint(255)
	myuint8 := uint8(255)
	mybyte := byte(255)
	myuint16 := uint16(65535)
	myuint32 := uint32(4294967295)
	myuint64 := uint64(18446744073709551615)
	myint8 := int8(-128)
	myint16 := int16(-32768)
	myint32 := int32(-2147483648)
	myrune := rune(-2147483648)
	myint64 := int64(-9223372036854775808)

	fmt.Println("myUncast - ", switchOnType(myUncast))
	fmt.Println("myint - ", switchOnType(myint))
	fmt.Println("myuint - ", switchOnType(myuint))
	fmt.Println("myuint8 - ", switchOnType(myuint8))
	fmt.Println("mybyte - ", switchOnType(mybyte))
	fmt.Println("myuint16 - ", switchOnType(myuint16))
	fmt.Println("myuint32 - ", switchOnType(myuint32))
	fmt.Println("myuint64 - ", switchOnType(myuint64))
	fmt.Println("myint8 - ", switchOnType(myint8))
	fmt.Println("myint16 - ", switchOnType(myint16))
	fmt.Println("myint32 - ", switchOnType(myint32))
	fmt.Println("myrune - ", switchOnType(myrune))
	fmt.Println("myint64 - ", switchOnType(myint64))
}

/*
Go is a statically typed programming language.
Variables always have a specific type and that type cannot change.

INTEGERS & FLOATING-POINT
Generally we split numbers into two different kinds: integers and floating-point numbers.

INTEGERS
uint8, uint16, uint32, uint64, int8, int16, int32 and int64
8, 16, 32 and 64 tell us how many bits each of the types use
uint means “unsigned integer” while int means “signed integer”
Generally if you are working with integers you should just use the int type.

UNSIGNED INTEGERS
only contain positive numbers (or zero).

ALIAS TYPES
byte which is the same as uint8
rune which is the same as int32

Bytes are an extremely common unit of measurement used on computers
(1 byte = 8 bits, 1024 bytes = 1 kilobyte, 1024 kilobytes = 1 megabyte, …)
and therefore Go's byte data type is often used in the definition of other types.

MACHINE DEPENDENT
uint, int and uintptr
machine dependent because their size depends on the type of architecture you are using.


uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

byte        alias for uint8
rune        alias for int32

https://golang.org/ref/spec#Numeric_types

*/
