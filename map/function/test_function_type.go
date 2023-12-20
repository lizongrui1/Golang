package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	type f1 func(int, int) int
	var ff f1
	ff = sum
	ff(1, 2)
	var r f1
	r = ff
	fmt.Printf("r: %v\n", r(1, 2))

	ff = max
	r = ff 
	result := r(3, 4)
	fmt.Printf("result: %v\n", result)
}
