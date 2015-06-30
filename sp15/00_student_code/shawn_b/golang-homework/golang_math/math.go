package math

import (
    "sort"
)
type byVal []float64
func (arr byVal) Len() int {return len(arr)}
func (arr byVal) Swap(i, j int) {arr[i], arr[j]=arr[j], arr[i]}
func (arr byVal) Less(i, j int) bool {return arr[i] < arr[j]}

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}

func Min(xs []float64) float64 {
    sort.Sort(byVal(xs))
    return (xs[0])
}

func Max(xs []float64) float64 {
    sort.Sort(byVal(xs))
    return (xs[len(xs)-1])
}
