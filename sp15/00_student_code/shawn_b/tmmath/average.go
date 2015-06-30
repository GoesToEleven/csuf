package tmmath

func Average(intSliceParam []int) int {

	total := 0
	count := 10

	for _, x := range intSliceParam {
		total += x
	}

	return total / count
}

// you must capitalize the func name for it to export (be available, be public)

func AverageCorrect(intSliceParam []int) int {

	total := 0
	count := 0

	for _, x := range intSliceParam {
		total += x
		count++
	}

	return total / count
}
