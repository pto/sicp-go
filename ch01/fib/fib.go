// Fib calculates Fibonacci numbers for the command line arguments.
package main

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	programName := filepath.Base(os.Args[0])
	if len(os.Args) < 2 || strings.HasPrefix(os.Args[1], "-h") {
		fmt.Fprintf(os.Stderr, "usage: %s <number>...\n", programName)
		os.Exit(1)
	}
	for _, arg := range os.Args[1:] {
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", programName, err)
			os.Exit(1)
		}
		fmt.Println(fib(n))
	}
}

// fib calculates the nth Fibonacci number.
func fib(n int) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	sum := new(big.Int)
	for i := 0; i < n; i++ {
		sum.Add(a, b)
		a.Set(b)
		b.Set(sum)
	}
	return a
}
