package sqrt

import (
	"math"
	"testing"
)

func Test_Sqrt(t *testing.T) {
	test_sqrt(2, t)
	test_sqrt(1e5, t)
	test_sqrt(1e-5, t)
	test_sqrt(1e50, t)
	test_sqrt(1e-50, t)
	test_sqrt(0, t)
	test_sqrt(-1, t)
	test_sqrt(-1e50, t)
	test_sqrt(-1e-50, t)
	test_sqrt(math.NaN(), t)
}

func test_sqrt(x float64, t *testing.T) {
	if math.IsNaN(math.Sqrt(x)) {
		if !math.IsNaN(Sqrt(x)) {
			t.Errorf("math.Sqrt(%v) is NaN, but Sqrt(%v) is not\n", x, x)
		}
	} else if math.Abs(Sqrt(x)-math.Sqrt(x)) > math.Sqrt(x)/1e-15 {
		t.Errorf("math.Sqrt(%v) is %v, but Sqrt(%v) is %v\n",
			x, math.Sqrt(x), x, Sqrt(x))
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
