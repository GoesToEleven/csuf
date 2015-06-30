package main

import (
	"fmt"
	"reflect"
	"runtime"
)


//Printer function 1
func myTextPrinter ( print_message string ) {
	fmt.Println(print_message)
}

//No printer available error message
func noPrinter( print_message string ) {
	print_message = "No printer available!"
	fmt.Println( print_message )
}

//Function factory, return the correct printer type
func getPrinter ( printType int) (func(string)) {
	switch( printType ) {
	case 0:
		return myTextPrinter
	default:
		return noPrinter
	}
}

func testPrinter(  printFunction func(string) ) {
		var fName = runtime.FuncForPC(reflect.ValueOf(printFunction).Pointer()).Name()
		fName = fName + " printer tested"
		printFunction( fName )
	}


func main() {

	var myMsg = "print this!"
	var printType = 0

	if printType == 0 {
		var aPrinter = getPrinter(0)
		aPrinter(myMsg)
	} else {
		fmt.Println("try again")
	}

	testPrinter( myTextPrinter )
}
