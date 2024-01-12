package main

import "fmt"

func test1() {
	fmt.Println("既没有参数又没有返回值")
}

func sum1(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func sum2(a int, b int) int {
	ret := a + b
	return ret
	//return a + b
}

func test2() (name string, age int) {
	/*name = "tom"
	age = 20
	return name, age*/

	n := "tom"
	a := 20
	return n, a
}

/*func test3() (string, int) {
	n := "tom"
	a := 20
	return n, a
}*/

func main() {
	//test1()
	//sum()
	name, age := test2()
	fmt.Printf("Name: %v, Age: %v\n", name, age)
}
