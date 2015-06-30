/*
Write a function with one variadic parameter that finds the greatest number in a list of numbers.
 */

package main

import (
	"fmt"
	"sort"
)

type byVal []int

func (arr byVal) Len() int {return len(arr)}
func (arr byVal) Swap(i, j int) {arr[i], arr[j]=arr[j], arr[i]}
func (arr byVal) Less(i, j int) bool {return arr[i] > arr[j]}

func findLarge( args ...int) int {
	sort.Sort(byVal(args))
	return args[0]
}

func main() {
	fmt.Println(findLarge(0, 1, 2, 10, 50, 20, 30))
}
