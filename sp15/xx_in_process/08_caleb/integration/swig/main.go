package main

import (
	"fmt"
	"math/rand"

	"github.com/calebdoxsey/tutorials/integration/swig/gsl"
)

func main() {
	xs := make([]float64, 4096)
	for i := 0; i < len(xs); i++ {
		xs[i] = rand.Float64()
	}
	fmt.Println(gsl.Mean(xs))
}
