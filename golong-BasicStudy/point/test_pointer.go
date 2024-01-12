package main

import "fmt"

func main() {
	var ip *int
	fmt.Printf("ip: %v\n", ip) //地址值
	fmt.Printf("ip: %T\n", ip) //类型

	var i int = 100
	ip = &i //取i的地址
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %v\n", *ip) //取这个地址对应的值

	var sp *string
	var s string = "tom"
	sp = &s //指针变量的存储地址
	fmt.Printf("sp: %v\n", sp)
	fmt.Printf("sp: %T\n", sp)
	fmt.Printf("sp: %T\n", *sp)

	var bp *bool
	var b bool = false
	fmt.Printf("b: %v\n", bp)
	fmt.Printf("b: %T\n", bp)
	bp = &b
	fmt.Printf("b: %v\n", *bp)``
}
