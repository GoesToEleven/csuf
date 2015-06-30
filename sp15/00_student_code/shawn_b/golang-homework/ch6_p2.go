/*
Write a program that finds the smallest number in this list:
x := []int{ 48,96,86,68, 57,82,63,70, 37,34,83,27, 19,97, 9,17, }
 */

package main

import (
	"fmt"
	"sort"
)

type byVal []int

func (arr byVal) Len() int {return len(arr)}
func (arr byVal) Swap(i, j int) {arr[i], arr[j]=arr[j], arr[i]}
func (arr byVal) Less(i, j int) bool {return arr[i] < arr[j]}

func main() {
	x := []int{ 48,96,86,68, 57,82,63,70, 37,34,83,27, 19,97, 9,17, }
	sort.Sort(byVal(x))
	fmt.Println(x[0])
}
