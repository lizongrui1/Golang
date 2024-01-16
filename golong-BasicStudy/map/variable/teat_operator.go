package main

import "fmt"

func getA() int{
	return 100
}

func main() {
	// a := 200
	// b := 50
	/* fmt.Printf("(a + b): %v\n", (a + b))
	fmt.Printf("(a - b): %v\n", (a - b)) */

	/* a++
	fmt.Printf("a: %v\n", a)
	b--
	fmt.Printf("b: %v\n", b) */

	/* fmt.Printf("(a == b): %v\n", (a == b))
	fmt.Printf("(a >= b): %v\n", (a >= b))
	fmt.Printf("(a <= b): %v\n", (a <= b)) */

	/* a := true
	b := false

	fmt.Printf("(a && b): %v\n", (a && b))
	fmt.Printf("(a || b): %v\n", (a || b))
	fmt.Printf("a: %v\n", !a) */

	/* a := 4
	fmt.Printf("a: %b\n", a)   //0100
	b := 8
	fmt.Printf("b: %b\n", b)   //1000

	r := a & b
	fmt.Printf("r: %v\n", r)   //0
	r = a | b
	fmt.Printf("r: %v\n", r)   //1100  -->  12     或
	r = a ^ b
	fmt.Printf("r: %v\n", r)   //1100  -->  12     异或
	r = a << 2
	fmt.Printf("r: %v\n", r)   //10000-->  16
	r = a >> 2
	fmt.Printf("r: %v\n", r)   //0001-->  1 */

	a := 100
	fmt.Printf("a: %v\n", a)
	a = 200
	fmt.Printf("a: %v\n", a)
	a = getA()
	fmt.Printf("a: %v\n", a)
	a += 100   //a = a + 100
	fmt.Printf("a: %v\n", a)


}
