// Package fact contains BigInt implementations of the factorial function.
package fact

import (
	"math/big"
)

var one = big.NewInt(1)

// FactRecursive calculates the factorial of x recursively.
func FactRecursive(x *big.Int) *big.Int {
	if x.Cmp(one) == 0 {
		return one
	}
	tmp := new(big.Int)
	return tmp.Mul(x, FactRecursive(tmp.Sub(x, one)))
}

// FactIterative calculates the factorial of x iteratively.
func FactIterative(x *big.Int) *big.Int {
	result := new(big.Int).Set(x)
	counter := new(big.Int).Sub(x, one)
	for counter.Cmp(one) > 0 {
		result.Mul(result, counter)
		counter.Sub(counter, one)
	}
	return result
}
