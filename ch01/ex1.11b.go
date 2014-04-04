// Triple-branching function, iterative
package main

import "fmt"

func main() {
	for i := 0; i <= 30; i++ {
		fmt.Printf("%2d %15d\n", i, f(i))
	}
}

func f(n int) int {
	a, b, c := 0, 1, 2
	for n > 0 {
		a, b, c = b, c, c+2*b+3*a
		n--
	}
	return a
}
