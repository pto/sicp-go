// Package fact contains big.Int implementations of the factorial function.
package fact

import (
	"math/big"
)

var one = big.NewInt(1)

// Recursive calculates the factorial of x recursively.
func Recursive(x *big.Int) *big.Int {
	if x.Cmp(one) == 0 {
		return big.NewInt(1)
	}
	r := Recursive(new(big.Int).Sub(x, one))
	return r.Mul(r, x)
}

// Iterative calculates the factorial of x iteratively.
func Iterative(x *big.Int) *big.Int {
	counter := new(big.Int).Sub(x, one)
	for counter.Cmp(one) > 0 {
		x.Mul(x, counter)
		counter.Sub(counter, one)
	}
	return x
}
