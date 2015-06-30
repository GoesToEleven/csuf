package min_max

import "testing"

type testpair struct {
	list      []float64
	maxresult float64
	minresult float64
}

var tests = []testpair{
	{[]float64{-5, 8, 19, 23, -9, 2}, 23, -9},
	{[]float64{0, 0, 0, 0, 0}, 0, 0},
	{[]float64{}, 0, 0},
}

func TestMin(t *testing.T) {
	for _, pair := range tests {
		i := Min(pair.list)
		if i != pair.minresult {
			t.Error(
				"For", pair.list,
				"expected", pair.minresult,
				"got", i)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range tests {
		i := Max(pair.list)
		if i != pair.maxresult {
			t.Error(
				"For", pair.list,
				"expected", pair.maxresult,
				"got", i)
		}
	}
}
