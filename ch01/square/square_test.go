package square

import (
	"math"
	"testing"
)

func TestSquare(t *testing.T) {
	cases := []float64{2, 1e5, 1e-5, 1e50, 1e-50, 0, 1 / math.Inf(-1),
		1 / math.Inf(1), -1, -1e50, -1e-50, math.NaN(), math.Inf(1),
		math.Inf(-1), math.SmallestNonzeroFloat64, math.MaxFloat64}

	for _, c := range cases {
		if (math.IsNaN(c*c) && !math.IsNaN(Square(c))) ||
			(math.IsNaN(Square(c)) && !math.IsNaN(c*c)) ||
			math.Abs(Square(c)-c*c) > c*c/1e-15 {
			t.Errorf("Square(%v) is %v, want %v\n", c, Square(c), c*c)
		}
	}
}

func Benchmark_Square(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Square(123456789.0)
	}
}

func Benchmark_MathSquare(b *testing.B) {
	square := func(x float64) float64 { return x * x }
	for i := 1; i < b.N; i++ {
		square(123456789.0)
	}
}
