package main

import "fmt"

// func main() {
//     var n int16 = 34
//     var m int32
//     m = int32(n)

//     fmt.Printf("32 bit int is: %d\n", m)
//     fmt.Printf("16 bit int is: %d\n", n)
// }

func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10 
	fmt.Printf("%b \n", a) // 1010 占位符%b表示二进制

	// 八进制 以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77
	// 十六进制 以8x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF
}