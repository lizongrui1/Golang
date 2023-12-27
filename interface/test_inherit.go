package main

import (
	"fmt"
)

type Animal struct { //定义一个父类
	name string
	age  int
}

func (a Animal) eat() { //父类的方法
	fmt.Println("eat...")
}

func (a Animal) sleep() { //父类的方法
	fmt.Printf("sleep...")
}

type Dog struct { //定义一个子类，嵌入父类
	a     Animal //可以理解为继承
	color string
}

type Cat struct { //定义一个子类，嵌入父类
	Animal
	gender string
}

func main() { //创建一个 dog 实例
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
