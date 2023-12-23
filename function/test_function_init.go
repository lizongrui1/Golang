package main

import "fmt"

var i int = initVar()

func init() {
	fmt.Println("init....") //先于main主函数执行的
}

func init(){
	fmt.Println("init2....")
}

func initVar() int {
	fmt.Println("initVar...")
	return 100
	
}

func main() {
	fmt.Println("main.....")
	
	/* 
	initVar
	init....
	init2....
	main..... */
