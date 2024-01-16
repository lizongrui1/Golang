package main

import "fmt"

func test1() {
	m1 := map[string]string{"name": "tom", "age": "20", "gender": "男"}
	for k := range m1 {
		fmt.Println(k)
	}
	fmt.Println("------------------")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	for _, v := range m1 {
		fmt.Println(v)
	}
}

func main() {
	test1()
}
