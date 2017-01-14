// Fixedpoint demonstrates fixed point algorithms
package main

import (
	"fmt"
	"math"
)

type floatFunc func(float64) float64

func closeEnough(v1, v2 float64) bool {
	const tolerance = 0.00001
	return math.Abs(v1-v2) < tolerance
}

func fixedpoint(f floatFunc, guess float64) float64 {
	for !closeEnough(guess, f(guess)) {
		guess = f(guess)
	}
	return guess
}

func averageDamp(f floatFunc) floatFunc {
	return func(x float64) float64 {
		return (x + f(x)) / 2
	}
}

func mySqrt(x float64) float64 {
	return fixedpoint(averageDamp(func(y float64) float64 {
		return x / y
	}), 1.0)
}

func main() {
	fmt.Println(fixedpoint(math.Cos, 1.0))
	fmt.Println(fixedpoint(func(y float64) float64 {
		return math.Sin(y) + math.Cos(y)
	}, 1.0))
	fmt.Println(mySqrt(2))
}
