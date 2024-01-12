package main

import "fmt"

/* func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
} */

func cal(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func main() {
	/* f := add()
	r := f(10)
	fmt.Printf("r: %v\n", r)
	r = f(20)
	fmt.Printf("r: %v\n", r)
	r = f(30)
	fmt.Printf("r: %v\n", r)
	fmt.Println("--------------")

	f1 := add()
	r = f1(100) //重新计数
	fmt.Printf("r: %v\n", r)
	r = f1(200)
	fmt.Printf("r: %v\n", r) */

	f1, f2 := cal(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
}
