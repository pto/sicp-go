package sqrt

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	cases := []float64{2, 1e5, 1e-5, 1e50, 1e-50, 0, 1 / math.Inf(-1),
		1 / math.Inf(1), -1, -1e50, -1e-50, math.NaN(), math.Inf(1),
		math.Inf(-1)}

	for _, c := range cases {
		if (math.IsNaN(math.Sqrt(c)) && !math.IsNaN(Sqrt(c))) ||
			(math.IsNaN(Sqrt(c)) && !math.IsNaN(math.Sqrt(c))) ||
			math.Abs(Sqrt(c)-math.Sqrt(c)) > math.Sqrt(c)/1e-15 {
			t.Errorf("Sqrt(%v) is %v, want %v\n", c, Sqrt(c), math.Sqrt(c))
		}
	}
}

func Benchmark_Sqrt(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Sqrt(123456789.0)
	}
}

func Benchmark_MathSqrt(b *testing.B) {
	for i := 1; i < b.N; i++ {
		math.Sqrt(123456789.0)
	}
}
