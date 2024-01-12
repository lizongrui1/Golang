package main

import "fmt"

func main() {
	a := [3]int{3, 6, 8}
	var pa [3]*int

	fmt.Printf("pa: %v\n", pa)

	for i := 0; i < len(a); i++ {
		pa[i] = &a[i]
	}
	fmt.Printf("pa: %v\n", pa)

	for i := 0; i < len(pa); i++ {
		fmt.Printf("pa[%v]: %v\n", i, *pa[i])
	}
}
