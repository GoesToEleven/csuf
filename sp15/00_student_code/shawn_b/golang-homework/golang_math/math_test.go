package math

import "testing"

type testpair struct {
	values []float64
	expected float64
}

var tests = []testpair{
	{ []float64{1,2}, 1.5 },
	{ []float64{1,1,1,1,1,1}, 1 },
	{ []float64{-1,1}, 0 },
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}

var testsMin = []testpair{
	{ []float64{1,2,3,4,5}, 1 },
	{ []float64{5,4,3,2,1}, 1 },
	{ []float64{5,4,0,1,2}, 0 },
	{ []float64{1,2,-2,0}, -2 },
}

func TestMin(t *testing.T) {
	for _, pair := range testsMin{
		v := Min(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}
var testsMax = []testpair{
	{ []float64{1,2,3,4,5}, 5 },
	{ []float64{5,4,3,2,1}, 5 },
	{ []float64{5,4,10,1,2}, 10 },
	{ []float64{-10,-20,-2,-30}, -2 },
}

func TestMax(t *testing.T) {
	for _, pair := range testsMax{
		v := Max(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}
