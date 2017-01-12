// Change counts the number of ways to make change.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var coinValues = []int{1, 5, 10, 25, 50}
var functionCalls = 0

func main() {
	programName := filepath.Base(os.Args[0])
	if len(os.Args) != 2 || strings.HasPrefix(os.Args[1], "-h") {
		fmt.Printf("usage: %s number-of-cents\n", programName)
		os.Exit(1)
	}
	amount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "%s:", programName, err)
		os.Exit(1)
	}
	ways := countChange(amount)
	fmt.Printf("%v ways in %v calls\n", ways, functionCalls)
}

func countChange(amount int) int {
	return cc(amount, len(coinValues))
}

func cc(amount int, coins int) int {
	functionCalls++
	switch {
	case amount == 0:
		return 1
	case amount < 0 || coins == 0:
		return 0
	default:
		return cc(amount, coins-1) + cc(amount-firstCoinValue(coins), coins)
	}
}

func firstCoinValue(coins int) int {
	return coinValues[coins-1]
}
