package main

import "fmt"

/*func test1() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("i：", i)
		} else {
			continue
		}
	}
}*/

func test2() {
	for i := 0; i < 5; i++ {
	LABEL:
		for j := 0; j < 4; j++ {
			if i == 2 && j == 2 {
				continue LABEL
			}
			fmt.Println(i, j)
		}
	}
}

func main() {
	//test1()
	test2()
}
