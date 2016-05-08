// Package ex08 calculates cube root using Newton's method
package ex08

import "math"

func Cbrt(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case math.IsNaN(x):
		return math.NaN()
	case math.IsInf(x, 1):
		return math.Inf(1)
	case math.IsInf(x, -1):
		return math.Inf(-1)
	}
	guess, prevGuess := 1.0, 0.0
	for !isGoodEnough(guess, prevGuess) {
		guess, prevGuess = improve(guess, x), guess
	}
	return guess
}

func isGoodEnough(guess, prevGuess float64) bool {
	return math.Abs(guess-prevGuess) < math.Abs(guess)/1e15
}

func improve(guess, x float64) float64 {
	return (x/(guess*guess) + 2*guess) / 3
}
