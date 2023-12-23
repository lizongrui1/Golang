package main

import (
	"fmt"
)

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() { //方法
	fmt.Println("eat...")
}

func (a Animal) sleep() {
	fmt.Printf("sleep...")
}

type Dog struct {
	a     Animal //可以理解为继承
	color string
}

type Cat struct {
	Animal
	gender string
}

func main() {
	dog := Dog{
		a:     Animal{name: "花花", age: 2},
		color: "褐色",
	}

	dog.a.eat()
	dog.a.sleep()
	fmt.Println()
	println(dog.color)
	println(dog.a.age)

	cat := Cat{
		Animal{name: "旺财", age: 3},
		"公",
	}
	cat.eat()
	cat.sleep()
	fmt.Println()
	println(cat.gender)

}
