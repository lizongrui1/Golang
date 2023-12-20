package main

import "fmt"

type Website struct{
	Name string
}


func main() {
	//  %v -->  var （变量）
	// fmt.Printf("\"hello\": %v\n", "hello")

	/* site := Website{Name:"lzzr"}
	fmt.Printf("site: %v\n", site)
	fmt.Printf("site: %#v\n", site)
	fmt.Printf("site: %T\n", site)
	
	a := true
	fmt.Printf("a: %t\n", a) */

	/* i := 96
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %b\n", i)   //代表二进制输出的5
	fmt.Printf("i: %c\n", i)   //代表ASCII对应的符号 */

	x := 100 
	p := &x   //指针
	fmt.Printf("p: %v\n", p)



}