package cbrt

import (
	"fmt"
	"math"
	"testing"
)

func TestCbrt(t *testing.T) {
	cases := []float64{2, 81, 1e5, 1e-5, 1e50, 1e-50, 0, 1 / math.Inf(-1),
		1 / math.Inf(1), -1, -1e50, -1e-50, math.NaN(), math.Inf(1),
		math.Inf(-1), math.SmallestNonzeroFloat64, math.MaxFloat64}

	for _, c := range cases {
		if (math.IsNaN(math.Cbrt(c)) && !math.IsNaN(Cbrt(c))) ||
			(math.IsNaN(Cbrt(c)) && !math.IsNaN(math.Cbrt(c))) ||
			math.Abs(Cbrt(c)-math.Cbrt(c)) > math.Abs(math.Cbrt(c)/1e-15) {
			t.Errorf("Cbrt(%v) is %v, want %v\n", c, Cbrt(c), math.Cbrt(c))
		}
	}
}

func Benchmark_Cbrt(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Cbrt(123456789.0)
	}
}

func Benchmark_MathCbrt(b *testing.B) {
	for i := 1; i < b.N; i++ {
		math.Cbrt(123456789.0)
	}
}

func ExampleCbrt() {
	iterations = 0
	result := Cbrt(2)
	fmt.Println(result, result*result*result, iterations)
	// Output:
	// 1.2599210498948732 2 6
}
