package sum

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var xs []int64

func init() {
	rand.Seed(0)
	xs = make([]int64, 4096)
	for i := 0; i < len(xs); i++ {
		xs[i] = int64(rand.Intn(2 << 30))
	}
}

func Test(t *testing.T) {
	assert := assert.New(t)

	type testCase struct {
		result int64
		arr    []int64
	}
	cases := []testCase{
		{0, []int64{}},
		{20, []int64{10, 10}},
		{4428505259927, xs},
	}

	for _, c := range cases {
		assert.Equal(c.result, SumV1(c.arr))
	}

	for _, c := range cases {
		assert.Equal(c.result, SumV2(c.arr))
	}

}

func BenchmarkV1(b *testing.B) {
	xs := make([]int64, 4096)
	for i := 0; i < len(xs); i++ {
		xs[i] = int64(rand.Intn(2 << 30))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumV1(xs)
	}
}

func BenchmarkV2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumV2(xs)
	}
}
