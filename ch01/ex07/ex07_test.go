package ex07_test

import (
	"math"
	"sicp/ch01/ex07"
	"testing"
)

var test_function = ex07.SqrtImproved

func Test_Original(t *testing.T) {
	f := ex07.Sqrt
	test_sqrt(f, 2, t)
	// test_sqrt(f, 2e-50, t) // FAILS
	// test_sqrt(f, 2e50, t)  // Stack overflow
}

func Test_Improved(t *testing.T) {
	f := ex07.SqrtImproved
	test_sqrt(f, 2, t)
	test_sqrt(f, 2e-50, t)
	test_sqrt(f, 2e50, t)
}

func test_sqrt(f func(float64) float64, x float64, t *testing.T) {
	sqrt := f(x)
	if math.Abs(sqrt*sqrt-x) > x/1e-15 {
		t.Errorf("%v(%v) squared is %v\n", f, x, sqrt*sqrt)
	}
}
