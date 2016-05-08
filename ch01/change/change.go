// Change counts the number of ways to make change.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const numberOfCoinDenominations = 5

var calls_to_cc int

func main() {
	if len(os.Args) != 2 || strings.HasPrefix(os.Args[1], "-h") {
		fmt.Printf("usage: %s number-of-cents\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	amount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "change:", err)
		os.Exit(1)
	}
	fmt.Printf("%v in %v calls\n", countChange(amount), calls_to_cc)
}

func countChange(amount int) int {
	return cc(amount, numberOfCoinDenominations)
}

func cc(amount int, coins int) int {
	calls_to_cc++
	switch {
	case amount == 0:
		return 1
	case amount < 0 || coins == 0:
		return 0
	default:
		return cc(amount, coins-1) + cc(amount-firstCoinSize(coins), coins)
	}
}

func firstCoinSize(coins int) int {
	m := map[int]int{1: 1, 2: 5, 3: 10, 4: 25, 5: 50}
	return m[coins]
}
