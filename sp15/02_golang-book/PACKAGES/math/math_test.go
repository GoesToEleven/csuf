package math

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

/*
The go test command will look for any tests in any of the files in the
current folder and run them. Tests are identified by starting a function
with the word Test and taking one argument of type *testing.T. In our case
since we're testing the Average function we name the test function TestAverage.

Once we have the testing function setup we write tests that use the code
we're testing. In this case we know the average of [1,2] should be 1.5
so that's what we check.
*/

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverageMas(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

/*
This is a very common way to setup tests (abundant examples can be found
in the source code for the packages included with Go). We create a struct
to represent the inputs and outputs for the function. Then we create a list
of these structs (pairs). Then we loop through each one and run the function.
*/
