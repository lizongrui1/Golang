package main

import "fmt"

func main() {
	type Person struct {
		id    int
		name  string
		age   int
		email string
	}

	/* var tom Person
	tom = Person{
		id:    10086,
		name:  "tom",
		age:   18,
		email: "12345@gmail.com",
	}
	*/
	/* var tom Person
	tom = Person{
		1,
		"tom",
		20,
		"12345@gmail.com",
	} */

	tom := Person{
		id : 10086,
		name: "tom",
	}

	fmt.Printf("tom: %v\n", tom)

}
