package main

import "fmt"
import "github.com/goestoeleven/golang-book/PACKAGES/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}

/*
Bundling code in packages serves 3 purposes:

It reduces the chance of having overlapping names.
This keeps our function names short and succinct

It organizes code so that its easier to find code you want to reuse.

It speeds up the compiler by only requiring recompilation of smaller chunks
of a program. Although we use the package fmt, we don't have to recompile it
every time we change our program.

-------------

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
