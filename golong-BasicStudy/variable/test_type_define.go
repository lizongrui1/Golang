package main

import "fmt"

func main() {
	/* type MyInt int       //类型定义
	var i MyInt
	i = 100
	fmt.Printf("i: %T,%v\n", i, i)    //i: main.MyInt,100    类型：main.MyInt   值：100 */

	type MyInt = int
	var i MyInt
	i = 100
	fmt.Printf("i: %T,%v\n", i, i)    //i: int,100

}
