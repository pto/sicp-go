// Triple-branching function, recursive
package main

import "fmt"

func main() {
	for i := 0; i <= 30; i++ {
		fmt.Printf("%2d %15d\n", i, f(i))
	}
}

func f(n int) int {
	if n < 3 {
		return n
	}
	return f(n-1) + 2*f(n-2) + 3*f(n-3)
}
