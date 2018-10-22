package main

import "os"

func Example() {
	os.Args = []string{"change", "100"}
	main()
	// Output:
	// 292 ways in 15499 calls
}
