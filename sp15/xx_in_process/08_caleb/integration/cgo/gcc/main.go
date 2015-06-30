package main

import "fmt"

//extern sum
func sum(*int64, int64) int64

func Sum(xs []int64) int64 {
	if len(xs) == 0 {
		return 0
	}
	return sum(&xs[0], int64(len(xs)))
}

func main() {
	fmt.Println(Sum([]int64{1, 78}))
}
