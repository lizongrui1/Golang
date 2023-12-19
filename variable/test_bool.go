package main

import "fmt"

func main() {
	/*var b1 bool = true
	var b2 = false
	b3 := true

	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("b3: %v\n", b3)*/

	//if
	/*age := 16
	if age >= 18 {
		fmt.Printf("你成年了！")
	} else {
		fmt.Printf("你还未成年！")
	}*/

	/*count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i = %v\n", i)
		//println(i)
	}*/

	age := 18
	gender := "男"
	if age >= 18 && gender == "男" {
		//fmt.Printf("\"你是成年男子\":%v\n", "你是成年男子")
		fmt.Printf("你是成年男子")
	}
}
