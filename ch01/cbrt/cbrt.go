// Package cbrt calculates cube roots by Newton's method.
package cbrt

import (
	"math"
)

var iterations int

// Cbrt returns the cube root of x.
func Cbrt(x float64) float64 {
	switch {
	case math.IsNaN(x):
		return math.NaN()
	case x == 0:
		return 0
	case math.IsInf(x, 1):
		return math.Inf(1)
	case math.IsInf(x, -1):
		return math.Inf(-1)
	}
	guess, previousGuess := 1.0, 0.0
	for !isGoodEnough(guess, previousGuess) {
		iterations++
		guess, previousGuess = improve(guess, x), guess
	}
	return guess
}

func isGoodEnough(guess float64, previousGuess float64) bool {
	return math.Abs(guess-previousGuess) < math.Abs(guess/1e15)
}

func improve(guess float64, x float64) float64 {
	return (x/(guess*guess) + 2*guess) / 3
}
