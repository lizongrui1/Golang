package main

import "fmt"

func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func comp(a int, b int) int {
	var max = 0
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	c := comp(1, 2)
	fmt.Println(c)
}
