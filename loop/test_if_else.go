package main

import "fmt"

func f1() {
	var num int
	fmt.Println("请输入一个数字：")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}

func f2() {
	if age := 20; age > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
}

func main() {
	/*a := 1
	b := 2
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}*/

	//f1()
	f2()
}
