package ex08

import (
	"math"
	"testing"
)

func Test_Cbrt(t *testing.T) {
	test_cbrt(-9, t)
	test_cbrt(2, t)
	test_cbrt(1e5, t)
	test_cbrt(1e-5, t)
	test_cbrt(1e50, t)
	test_cbrt(1e-50, t)
	test_cbrt(0, t)
	test_cbrt(1/math.Inf(-1), t) // negative zero
	test_cbrt(-9, t)
	test_cbrt(-1, t)
	test_cbrt(-1e50, t)
	test_cbrt(-1e-50, t)
	test_cbrt(math.NaN(), t)
	test_cbrt(math.Inf(1), t)
	test_cbrt(math.Inf(-1), t)
}

func test_cbrt(x float64, t *testing.T) {
	if (math.IsNaN(math.Cbrt(x)) && !math.IsNaN(Cbrt(x))) ||
		(math.IsNaN(Cbrt(x)) && !math.IsNaN(math.Cbrt(x))) ||
		math.Abs(Cbrt(x)-math.Cbrt(x)) > math.Abs(math.Cbrt(x))/1e-15 {
		t.Errorf("Cbrt(%v) is %v, not %v\n", x, Cbrt(x), math.Cbrt(x))
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
