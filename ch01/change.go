// Count number of ways to make change for a dollar
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type cents int
type coinCount int
type changeCount int

var calls_to_cc int

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	var (
		amount int
		err    error
	)
	if amount, err = strconv.Atoi(os.Args[1]); err != nil {
		usage()
	}
	fmt.Printf("%v in %v calls\n", countChange(cents(amount)), calls_to_cc)
}

func usage() {
	fmt.Printf("usage: %s coint-count\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func countChange(amount cents) changeCount {
	return cc(amount, 5)
}

func cc(amount cents, coins coinCount) changeCount {
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

func firstCoinSize(coins coinCount) cents {
	m := map[coinCount]cents{1: 1, 2: 5, 3: 10, 4: 25, 5: 50}
	return m[coins]
}
