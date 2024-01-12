package main

import "fmt"

/*func f1() {
	a, b, c := 1, 21, 3
	if a > b {
		if a > c {
			fmt.Println("a最大")
		}
	} else {
		// b > a
		if b > c {
			fmt.Println("b最大")
		} else {
			fmt.Println("c最大")
		}
	}
}*/

func f3() {
	gender := "男"
	age := 20
	if gender == "男" {
		if age >= 18 {
			fmt.Println("成年男性")
		} else {
			fmt.Println("未成年男性")
		}
	} else {
		if age >= 18 {
			fmt.Println("成年女性")
		} else {
			fmt.Println("未成年女性")
		}
	}
}

func main() {
	//f1()
	f2()
}
