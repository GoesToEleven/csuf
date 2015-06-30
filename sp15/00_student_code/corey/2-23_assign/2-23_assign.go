package main

import (
	"fmt"

	"github.com/saxleader/golang/2-23_assign/slice_methods"
)

type Name struct {
	first MyStr
	last  MyStr
}

type MyStr string

func (s MyStr) Initial() string {
	return string(s[0])
}

func (n Name) FullName() MyStr {
	return n.first + " " + n.last
}

func PrintSlice(s []string) {
	for _, current := range s {
		fmt.Println(current)
	}
}

func main() {
	pies := slice_methods.StrSlice{"Apple", "Peach", "Cherry", "Blueberry", "Chocolate", "Pecan"}
	pieArray := [6]string{"Apple", "Peach", "Cherry", "Blueberry", "Chocolate", "Pecan"}
	same := true
	for key, current := range pies {
		if current == pieArray[key] {
			continue
		} else {
			same = false
			fmt.Println("The slice and Array are not equal.")
		}
	}
	if same {
		fmt.Println("The Slice and Array are equal.")
	}
	fmt.Println("\nAvailable Flavors of Pie:")
	PrintSlice(pies)

	fmt.Println("\nFirst Pie:")
	pies.PrintFirst()

	secondHalf := pies.SecondHalf()
	fmt.Println("\nSecond Half of Menu:")
	PrintSlice(secondHalf)

	fruit := pies[:2]
	fruit = append(fruit, pies[2:4]...)
	fmt.Println("\nFruit Pies:")
	PrintSlice(fruit)

	noChocolate := pies[:]
	noChocolate = append(noChocolate[:4], noChocolate[5:]...)
	fmt.Println("\nPies excluding Chocolate:")
	PrintSlice(noChocolate)

	myname := Name{"Corey", "Dihel"}
	fmt.Printf("\nMy name is %s\n", myname.FullName())
	fmt.Printf("My initials are %s", myname.first.Initial()+myname.last.Initial())
}
