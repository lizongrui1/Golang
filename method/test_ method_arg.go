package main

import (
	"fmt"
)

func test1() {
	type Person struct {
		name string
	}

	p1 := Person{
		name: "tom",
	}

	p2 := &Person{
		name: "tom",
	}

	fmt.Printf("p1: %T\n", p1)
	fmt.Printf("p2: %T\n", p2)
}

type Person struct {
	name string
}

/* func showPerson1(per Person) {
	per.name = "kite"
}

func showPerson2(per *Person) {
	//per.name （自动解引用）
	//正常都是（*per）.name = “tom”
	per.name = "kite"
} */

func (per Person) showPerson3() {
	per.name = "kite"
}

func (per *Person) showPerson4() {
	//per.name （自动解引用）
	//正常都是（*per）.name = “tom”
	per.name = "kite"
}

func main() {
	p1 := Person{
		name: "tom",
	}

	p2 := &Person{
		name: "tom",
	}

	/* showPerson1(p1)
	fmt.Printf("p1: %v\n", p1)      //p1: {tom}
	fmt.Println("---------------")
	showPerson2(p2)
	fmt.Printf("p2: %v\n", *p2)     //p2: {kite} */

	p1.showPerson3()
	fmt.Printf("p1: %v\n", p1)
	fmt.Println("----------------")
	p2.showPerson4()
	fmt.Printf("p2: %v\n", *p2)
}
