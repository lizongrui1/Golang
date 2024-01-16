package main

import "fmt"

func f1() {
	var score int
	fmt.Println("请输入考试分数")
	fmt.Scan(&score)

	if score >= 60 && score <= 70 {
		fmt.Println("C")
	} else if score >= 70 && score <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}

func main() {
	f1()
}
