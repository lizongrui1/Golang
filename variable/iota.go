package main

import "fmt"

func main() {
	//iota
	/*	const (
			a1 = iota
			a2 = iota
			a3 = iota
		)

		fmt.Printf("a1: %v\na2: %v\na3: %v\n", a1, a2, a3)
		//fmt.Printf("a2: %v\n", a2)
		//fmt.Printf("a3: %v\n", a3)*/

	/*	const (
			a1 = iota
			_
			_
			a2 = iota
		)
		fmt.Printf("a1： %v\n", a1)
		fmt.Printf("a2： %v\n", a2)*/

	const (
		a1 = iota
		a2 = 100
		a3 = iota
	)
	fmt.Printf("a1： %v\n", a1)
	fmt.Printf("a2： %v\n", a2)
	fmt.Printf("a3： %v\n", a3)
}
