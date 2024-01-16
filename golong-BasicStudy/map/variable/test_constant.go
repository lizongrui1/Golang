package main

import "fmt"

func main() {
	const PI float32 = 3.14
	const PI2 = 3.1415
	fmt.Printf("PI:%v\n", PI)

	const (
		A = 100
		B = 200
	)
	fmt.Printf("A: %v\n", A)
	fmt.Printf("B: %v\n", B)

	const w, h = 200, 300
	fmt.Printf("w: %v\n", w)
	fmt.Printf("h: %v\n", h)
}
