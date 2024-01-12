package main

import (
	"fmt"
	main2 "study/src/self-study/main"
)

/*func test1() {
	i := 1
	if i >= 2 {
		fmt.Println("2")
	} else {
		goto END
	}
END:
	fmt.Printf("END...")
}*/

func test2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i >= 2 && j >= 2{
			goto END
		}
		fmt.Printf(i,j)
	}
END:
	fmt.Println("END.")
}

func main2.main() {
	//test1()
	test2()
}