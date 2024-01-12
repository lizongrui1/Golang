package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func f1(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func f2(name string, age int, ok bool, args ...int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("ok: %v\n", ok)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}

}

func main() {
	/* r := sum()
	fmt.Printf("r: %v\n", r) */
	f1(1, 2, 3, 4)
	f1(2, 3, 4, 7)
	f2("tom",20,true,12,3,9)

}
