// Package sqrt calculates square roots by Newton's method.
package sqrt

import "math"

func Sqrt(x float64) float64 {
	switch {
	case x < 0 || math.IsNaN(x):
		return math.NaN()
	case x == 0:
		return 0
	case math.IsInf(x, 1):
		return math.Inf(1)
	}
	guess, previousGuess := 1.0, 0.0
	for !isGoodEnough(guess, previousGuess) {
		guess, previousGuess = improve(guess, x), guess
	}
	return guess
}

func isGoodEnough(guess float64, previousGuess float64) bool {
	return math.Abs(guess-previousGuess) < guess/1e15
}

func improve(guess float64, x float64) float64 {
	return (guess + x/guess) / 2
}