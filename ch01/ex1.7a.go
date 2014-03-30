// Show that the original square root function fails on
// very large and very small numbers.
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(2e-50))
	fmt.Println(sqrt(2e50))
}

func sqrt(x float64) float64 {
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
