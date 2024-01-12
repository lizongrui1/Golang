package main

import "fmt"

/*func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i >= 5 {
			break
		}
	}
}*/

/*func test2() {
	i := 2
	switch i {
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Printf("2")
		break
		fallthrough //3不会被打印
	case 3:
		fmt.Printf("3")

	}
}*/

func test3() {
MYLABEL:
	for i := 0; i < 10; i++ {
		if i >= 5 {
			break MYLABEL
		}
	}
	fmt.Println("END...")
}

func main() {
	//test1()
	//test2()
	test3()
}
