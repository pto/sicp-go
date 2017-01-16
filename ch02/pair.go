package main

import "fmt"

type Pair struct {
	Car interface{}
	Cdr *Pair
}

func Cons(i interface{}, p *Pair) *Pair {
	return &Pair{i, p}
}

func (p *Pair) String() string {
	var result = "["
	var s string
	for p != nil {
		result += s + fmt.Sprint(p.Car)
		s = ", "
		p = p.Cdr
	}
	return result + "]"
}

func main() {
	list := Cons(1, Cons(2, Cons(3, nil)))
	fmt.Println(list)

	list = Cons("a", Cons("b", Cons("c", Cons("d", nil))))
	fmt.Println(list)

	list = Cons(1.1, Cons(2.2, Cons(3.3, Cons(4.4, Cons(5.5, nil)))))
	fmt.Println(list)

	list = nil
	fmt.Println(list)
}
