package main

import (
	"fmt"
)

type easiness struct {
	scale int
	desc string
}

type controls struct {
	manufacturer string
	cost int
	usability easiness
}

const (
	const_easy = 1
	const_medium = 2
	const_hard  = 3
)

func getSimpleBrief( arg_controls controls)(string,string) {
	return arg_controls.manufacturer, arg_controls.usability.desc
}

func main() {

	var ControlLogix = controls{
		manufacturer: "Allen-Bradley",
		cost: 10000,
		usability: easiness{
			scale: const_easy,
			desc: "easy",
		},
	}

	//fmt.Println("Hello World")

	var man, _ = getSimpleBrief( ControlLogix )
	fmt.Println(man)
	fmt.Println(ControlLogix.usability.scale)
	fmt.Println(const_easy)
}
