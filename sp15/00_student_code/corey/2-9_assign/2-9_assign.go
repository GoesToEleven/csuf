package main

import (
	"fmt"
)

const (
	A = iota
	B = iota
	C = iota
	D = iota
	E = iota
	F = iota
	G = iota
)

type Message struct {
	author string
	body   string
}

func my_func(new_message Message) {
	fmt.Println(new_message.author, "says:", new_message.body)
}

func multi_return(nums ...int) (int, int, int) {
	size := len(nums)
	return nums[size-3], nums[size-2], nums[size-1]
}

func main() {
	var message1 = Message{"Corey", "Hello"}
	var message2 = Message{"Taylor", "Hi"}
	var message3 = Message{"Megan", "Nice to meet you"}
	_, sec_to_last, last := multi_return(A, B, C, D, E, F, G)
	third_to_last, _, _ := multi_return(A, B, C, D, E, F, G)
	fmt.Println("Third to last number is", third_to_last)
	fmt.Println("Second to last number is", sec_to_last)
	fmt.Println("Last number is", last, "\n")
	my_func(message1)
	my_func(message2)
	my_func(message3)
}
