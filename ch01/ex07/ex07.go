// Package ex07 implements square roots by Newton's method
package ex07

import (
	"math"
)

// Original square root function from text.
func Sqrt(x float64) float64 {
	return sqrtIter(1, x)
}

func sqrtIter(guess, x float64) float64 {
	if isGoodEnough(guess, x) {
		return guess
	}
	return sqrtIter(improvedGuess(guess, x), x)
}

func isGoodEnough(guess, x float64) bool {
	return math.Abs(guess*guess-x) < 1e-15
}

func improvedGuess(guess, x float64) float64 {
	return average(guess, x/guess)
}

func average(a, b float64) float64 {
	return (a + b) / 2
}

// Improved square root function that works for very large
// and very small numbers.
func SqrtImproved(x float64) float64 {
	return sqrtIterImproved(1, 0, x)
}

func sqrtIterImproved(guess, prevGuess, x float64) float64 {
	if isGoodEnoughImproved(guess, prevGuess) {
		return guess
	}
	return sqrtIterImproved(improvedGuess(guess, x), guess, x)
}

func isGoodEnoughImproved(guess, prevGuess float64) bool {
	return math.Abs(guess-prevGuess)/guess < 1e-15
}
