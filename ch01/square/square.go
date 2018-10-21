// Package square calculates squares with logarithms.
package square

import "math"

// Square returns the square of x.
func Square(x float64) float64 {
	if x == 0 {
		return 0
	}
	if x < 0 {
		x = -x
	}
	return math.Exp(2 * math.Log(x))
}
