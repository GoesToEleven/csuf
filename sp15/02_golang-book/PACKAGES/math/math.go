package math

// Average finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

/*
Using a terminal in the math folder you just created run go install.
This will compile the math.go program and create a linkable object file:
workspace/pkg/darwin_amd64/github.com/goestoeleven/math.a

In Go if something starts with a capital letter that means other packages
(and programs) are able to see it. If we had named the function average
instead of Average our main program would not have been able to see it.

It's a good practice to only expose the parts of our package that we want
other packages using and hide everything else. This allows us to freely change
those parts later without having to worry about breaking other programs,
and it makes our package easier to use.

Package names match the folders they fall in. There are ways around this,
but it's a lot easier if you stay within this pattern.
*/
