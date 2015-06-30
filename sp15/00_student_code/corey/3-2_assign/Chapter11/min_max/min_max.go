package min_max

// Finds the minimum value in a given slice of float64
func Min(s []float64) float64 {
	min := s[0]
	for _, value := range s {
		if value < min {
			min = value
		}
	}
	return min
}

// Finds the maximum value in a given slice of float64
func Max(s []float64) float64 {
	max := s[0]
	for _, value := range s {
		if value > max {
			max = value
		}
	}
	return max
}
