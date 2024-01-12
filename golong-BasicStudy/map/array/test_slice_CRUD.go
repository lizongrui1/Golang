package main

import "fmt"

//增删改查 CRUD

// add
/*func test1() {
	var s1 = []int{}
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	s1 = append(s1, 300)
	fmt.Printf("s1:%v\n", s1)
}*/

// delete
/*func test2() {
	var s1 = []int{1, 2, 3, 4}
	//想删除i为2的数字
	s1 = append(s1[:2], s1[3:]...)
	//a = append(a[:index], a[index+1:]...)
	fmt.Printf("%v\n", s1)
}*/

// update
/*func test3() {
	var s1 = []int{1, 2, 3, 4}
	s1[1] = 100
	fmt.Printf("%v\n", s1)
}*/

// query
/*func test4() {
	var s1 = []int{1, 2, 3, 4}
	var key = 2
	for i, v := range s1 {
		if v == key {
			fmt.Printf("%v\n", i) //2的索引是1
		}
	}
}*/

// copy
func test5() {
	var s1 = []int{1, 2, 3, 4}
	//var s2 = s1
	var s2 = make([]int, 4)
	copy(s2, s1)
	fmt.Printf("%v\n", s1)
	s2[0] = 100
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	test5()
}
