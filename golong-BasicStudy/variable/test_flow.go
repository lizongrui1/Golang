package main

import "fmt"

func main() {
	fmt.Println("aaa")
	fmt.Println("bbb")

	fmt.Printf("--------------\n")
	
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	fmt.Printf("--------------\n")

	x := [...]int{1, 2, 3}
	for _, v := range x {
		fmt.Printf("v: %v\n", v)
	}
	
}
