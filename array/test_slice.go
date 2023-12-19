package main

/*func test1() {
	//var a1 = []int{1, 2, 3}
	//fmt.Printf("%v\n", a1)

	var s1 []int
	var s2 []string
	fmt.Printf("s1：%v\n", s1)
	fmt.Printf("s2：%v\n", s2)
}*/

/*func test2() {
	//make
	var s1 = make([]int, 2)
	fmt.Printf("%v\n", s1)
}*/

func test3() {
	var s1 = []int{1, 2, 3}
	println(len(s1))
	println(cap(s1))

	println("s1:", s1[1])
}

func main() {
	//test1()
	//test2()
	test3()
}
