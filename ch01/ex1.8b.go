package main

import (
	"fmt"
	"math"
)

func main() {
	printResult(2)
	printResult(2e-50)
	printResult(2e50)
}

func printResult(x float64) {
	root := cubeRoot(x)
	fmt.Printf("%25.15g %25.15g\n", root, root*root*root)
}

func cubeRoot(x float64) float64 {
	guess, prevGuess := 1.0, 0.0
	for !isGoodEnough(guess, prevGuess) {
		guess, prevGuess = improvedGuess(guess, x), guess
	}
	return guess
}

func isGoodEnough(guess, prevGuess float64) bool {
	return math.Abs(guess-prevGuess)/guess < 1e-15
}

func improvedGuess(guess, x float64) float64 {
	return (x/(guess*guess) + 2*guess) / 3
}
