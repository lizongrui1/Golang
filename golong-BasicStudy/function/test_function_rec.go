package main

import "fmt"

/* func f1() {
	var s int = 1
	for i := 1; i <= 5; i++ {          //5的阶乘
		s = s * i
	}
	fmt.Printf("s: %v\n", s)
} */

func f2(a int) int {
	if a == 1 {
		return 1
	} else {
		return a * f2(a-1)
	}
}

func main() {
	// f1()

	// i := f2(5)
	// fmt.Printf("i: %v\n", i)

	fmt.Printf("i:%v",f2(5))
}
