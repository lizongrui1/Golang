package main

import "fmt"

/*func test1() {
	var s1 = []int{1, 2, 3, 4}
	l := len(s1)
	for i := 0; i < l; i++ {
		//println(s1[i])
		fmt.Printf("s1[%v]: %v\n", i, s1[i])
	}
}*/

func test2() {
	var s1 = []int{1, 2, 3, 4}
	//for i, v := range s1 {
	//	fmt.Printf("s1[%v]: %v\n", i, v)
	//}

	for _, v := range s1 {
		fmt.Printf("%v\n", v)
	}
}

func main() {
	//test1()
	test2()
}
