// Pascal's triangle
package main

import (
	"fmt"
	"strings"
)

const LINES = 10

func main() {
	for i := 1; i <= LINES; i++ {
		fmt.Print(strings.Repeat("  ", LINES-i))
		for _, v := range triangle_line(i) {
			fmt.Printf("%4d", v)
		}
		fmt.Println()
	}
}

func triangle_line(n int) []int {
	if n == 1 {
		return []int{1}
	}
	prev := triangle_line(n - 1)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			result[i] = 1
		} else {
			result[i] = prev[i] + prev[i-1]
		}
	}
	return result
}
