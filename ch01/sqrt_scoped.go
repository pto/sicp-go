package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(2))
}

func sqrt(x float64) float64 {
	isGoodEnough := func(guess float64) bool {
		return math.Abs(guess*guess-x) < 1e-15
	}
	improvedGuess := func(guess float64) float64 {
		return (guess + x/guess) / 2
	}
	var sqrtIter func(float64) float64
	sqrtIter = func(guess float64) float64 {
		if isGoodEnough(guess) {
			return guess
		}
		return sqrtIter(improvedGuess(guess))
	}
	return sqrtIter(1)
}
