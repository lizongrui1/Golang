package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (per Person) eat() {
	fmt.Println("eat。。。")
}

func (per Person) sleep() {
	fmt.Printf("sleep...")
}

func (per Person) work() {
	fmt.Printf("work...")
}

func main() {
	per := Person{
		name: "tom",
		age:  18,
	}

	fmt.Printf("%v\n", per)
	per.eat()
	per.sleep()
	per.work()
}
