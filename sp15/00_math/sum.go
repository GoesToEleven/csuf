package math

func Sum(intSliceParam []int) int {

	total := 0

	for _, x := range intSliceParam {
		total += x
	}

	return total
}
