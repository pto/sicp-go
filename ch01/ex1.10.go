// Ackermann's function
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}
	var arg1, arg2 int
	var err error
	if arg1, err = strconv.Atoi(os.Args[1]); err != nil {
		usage()
	}
	if arg2, err = strconv.Atoi(os.Args[2]); err != nil {
		usage()
	}
	fmt.Println(a(arg1, arg2))
}

func usage() {
	fmt.Printf("usage: %v arg1 arg2\n", filepath.Base(os.Args[0]))
	os.Exit(-1)
}

func a(x, y int) int {
	switch {
	case y == 0:
		return 0
	case x == 0:
		return 2 * y
	case y == 1:
		return 2
	default:
		return a(x-1, a(x, y-1))
	}
}
