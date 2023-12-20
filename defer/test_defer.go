package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("step1")
	defer fmt.Println("step2")
	fmt.Println("end")
}