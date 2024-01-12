package main

import "fmt"

/*func test1() {
	var a1 [2]int
	var a2 [3]string
	var a3 [1]bool

	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)

	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", a3)
}*/

/*func test2() {
	var a1 = [4]int{1, 2}
	fmt.Printf("%v\n", a1)
	var a2 = [2]string{"hello", "world"}
	fmt.Printf("%v\n", a2)
	var a3 = [2]bool{true, false}
	fmt.Printf("%v\n", a3)
}*/

/*func test3() {
	//数组的初始化，省略长度
	var a1 = [...]int{1, 2, 3, 4, 5}
	var a2 = [...]string{"ni", "hao"}
	var a3 = [...]bool{false, true}
	fmt.Printf("len(a1):%v\n", len(a1))
	fmt.Printf("len(a2):%v\n", len(a2))
	fmt.Printf("len(a3):%v\n", len(a3))
}*/

/*func test4() {
	var a1 = [...]int{1: 3, 3: 68, 5: 100}
	var a2 = [...]string{0: "hhh"}
	var a3 = [...]bool{9: true}
	fmt.Printf("a1:%v\n", a1)
	fmt.Printf("a2:%v\n", a2)
	fmt.Printf("a2:%v\n", a3)
}*/

func test5() {
	var a1 = [...]int{1, 2, 5}
	fmt.Printf("a1：%v\n", a1)
	a1[2] = 3
	fmt.Printf("a1: %v\n", a1)
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	test5()
}
