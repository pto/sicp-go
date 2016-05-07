// Package ex03 is a solution to SICP exercise 1-3.
package ex03

// Return the sum of squares of the two largest of three numbers
func SumOfSquaresLargestTwo(a, b, c float64) float64 {
	switch {
	case a <= b && a <= c:
		return sumOfSquares(b, c)
	case b <= a && b <= c:
		return sumOfSquares(a, c)
	default:
		return sumOfSquares(a, b)
	}
}

func sumOfSquares(x, y float64) float64 {
	return x*x + y*y
}
