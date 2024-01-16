package main

import "fmt"

func test1() {
	var name string
	name = "tom"
	var p_name *string
	p_name = &name

	fmt.Printf("name: %v\n", name)
	fmt.Printf("p_name: %v\n", p_name)  //指针地址
	fmt.Printf("p_name: %v\n", *p_name) //根据地址取值
}

func test2() {
	type Person struct {
		id   int
		name string
		age  int
	}

	tom := Person{
		id:   10086,
		name: "tom",
		age:  18,
	}

	var p_person *Person
	p_person = &tom
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %p\n", p_person)
	fmt.Printf("p_person: %v\n", *p_person)
}

func test3() {
	type Person struct {
		id   int
		name string
		age  int
	}

	var tom = new(Person)
	tom.id = 10086
	tom.name = "tom"
	tom.age = 18
	fmt.Printf("tom: %v\n", *tom)
}

func main() {
	test3()
}
