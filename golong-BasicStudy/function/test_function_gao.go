package main

import (
	"fmt"
)

func sayHello(name string) {
	fmt.Printf("Hello: %v\n", name)
}

func test(name string, f func(string)) {
	f(name)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func cal(operator string) func(int, int) int {
	switch operator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	// test("tom", sayHello)

	ff := cal("+")
	r := ff(10,8)
	fmt.Printf("r: %v\n", r)

	ff = cal("-")
	r = ff(14,0)
	fmt.Printf("r: %v\n", r)

}
