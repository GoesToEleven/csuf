package sum

func SumV1(xs []int64) int64 {
	var total int64
	for i := 0; i < len(xs); i++ {
		total += xs[i]
	}
	return total
}

func SumV2(xs []int64) int64
