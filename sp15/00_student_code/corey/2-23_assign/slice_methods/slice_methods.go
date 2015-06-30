package slice_methods

import (
	"fmt"
)

type StrSlice []string

func (s StrSlice) PrintFirst() {
	fmt.Println(s[0])
}

func (s StrSlice) SecondHalf() StrSlice {
	return s[(len(s) / 2):]
}
