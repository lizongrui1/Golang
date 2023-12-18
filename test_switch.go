/*package main

import (
	"flag"
	"fmt"
)

func f4() {
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("一般")
	}
}

func f5() {
	day := 3
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	default:
		fmt.Println("非法输入")
	}
}

func f6() {
	var score int
	fmt.Println("请输入分数")
	fmt.Scan(score)

	switch  {
	case score >= 60:
		fmt.Println("及格")
	case score >= 80 && score <= 100:
		fmt.Println("优秀")

	}


func f7() {
	num := 100
		switch num {
		case 100:
			fmt.Println("100")
			fallthrough
		case 200:
			fmt.Println("200")
		case 300:
			fmt.Println("300")
		}
	}

func main(){
	//f4()
	//f5()
	//f6()
	f7()
}
*/