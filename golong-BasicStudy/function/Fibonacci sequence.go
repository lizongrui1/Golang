package main

import "fmt"

func f(n int) int {
	if n == 1 || n == 2 {
		return 1
	} /* else {
		return f(n-1) + f(n-2)
	} */
	return f(n-1) + f(n-2)

}

func main() {
	r := f(5)
	fmt.Printf("r: %v\n", r)
}