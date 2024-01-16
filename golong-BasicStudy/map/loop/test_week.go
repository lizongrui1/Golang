package main

import "fmt"

func week() {
	var day string
	fmt.Println("请输入一个字符")
	fmt.Scan(&day)

	//判断Monday Tuesday Wednesday Thursday Friday Saturday Sunday
	if day == "T" {
		fmt.Println("请再输入下一个字符")
		fmt.Scan(&day)
		if day == "u" {
			fmt.Println("是星期二")
		} else {
			fmt.Println("是星期四")
		}
	} else if day == "S" {
		fmt.Println("请再输入下一个字符")
		fmt.Scan(&day)
		if day == "a" {
			fmt.Println("是星期六")
		} else {
			fmt.Println("是星期日")
		}
	} else if day == "M" {
		fmt.Println("是星期一")
	} else if day == "W" {
		fmt.Println("是星期三")
	} else {
		fmt.Println("是星期五")
	}
}

func main() {
	week()
}
