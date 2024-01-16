package main

import "fmt"

type Person struct {
	id   int
	name string
}

func showPerson(per Person) {
	per.id = 10086
	per.name = "kit"
	fmt.Printf("per: %v\n", per)
}

func showPerson2(per *Person) {
	per.id = 10087
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

//值传递拷贝了一份副本
func main() {
	tom := Person{
		id : 10000,
		name : "tom",
	}
	/* fmt.Printf("tom: %v\n", tom)
	fmt.Println("-----------------")

	showPerson(tom)
	fmt.Printf("tom: %v\n", tom) */

	per := &tom
	fmt.Printf("tom: %v\n", tom)
	fmt.Println("-----------------")
	showPerson2(per)
	fmt.Printf("per: %v\n", per)
	
}