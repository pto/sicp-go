// Iterative addition
package main

import "fmt"

type value struct {
	integer int
	text    string
}

func main() {
	result := add(value{4, "4"}, value{5, "5"})
	fmt.Println(result.integer)
	fmt.Println(result.text)
}

func inc(a value) value {
	return value{a.integer + 1, fmt.Sprintf("inc(%v)", a.text)}
}

func dec(a value) value {
	return value{a.integer - 1, fmt.Sprintf("dec(%v)", a.text)}
}

func add(a, b value) value {
	if a.integer == 0 {
		return value{b.integer, fmt.Sprintf("add(%v, %v)", a.text, b.text)}
	}
	return add(dec(a), inc(b))
}
