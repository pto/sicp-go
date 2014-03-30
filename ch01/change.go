// Count number of ways to make change for a dollar
package main

import "fmt"

type cents int
type coinCount int
type changeCount int

func main() {
	fmt.Println(countChange(100))
}

func countChange(amount cents) changeCount {
	return cc(amount, 5)
}

func cc(amount cents, coins coinCount) changeCount {
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
	switch coins {
	case 1:
		return 1
	case 2:
		return 5
	case 3:
		return 10
	case 4:
		return 25
	case 5:
		return 50
	}
	panic("wrong number of coins")
}
