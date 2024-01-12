package main

/*func test1() {
	var a1 [2]int
	a1[0] = 100
	a1[1] = 200
	fmt.Printf("%v\n", a1[0])
	fmt.Printf("%v\n", a1[1])

	fmt.Println("------------------")
	a1[0] = 10
	a1[1] = 20
	fmt.Printf("%v\n", a1[0])
	fmt.Printf("%v\n", a1[1])

	//数组长度越界
	//a1[3] = 1       错误

}*/

/*func test2() {
	var a1 = [3]int{1, 2, 3}
	//println(len(a1))
	fmt.Printf("%v\n", len(a1))

	var a2 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", len(a2))
}*/

/*func test3() {
	//数组的遍历1        根据长度和小标
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(a); i++ {
		//fmt.Printf("i：%v\n", a[i])
		println(a[i])
	}

}*/

func test4() {
	//数组的遍历2    for range
	var a = [...]int{1, 2, 3, 4}
	for _, value := range a {
		println(value)
	}
}

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}
