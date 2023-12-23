package main

import "fmt"

type Person struct {
	name string
}

func (per Person) eat() {
	fmt.Printf("per.name: %v, eat...\n", per.name)
}

func (per Person) sleep() {
	fmt.Printf("per.name: %v, sleep...\n", per.name)
}

type Customer struct {
	name string
}

func (customer Customer) login(name string, password string) bool {
	fmt.Printf("name:%v\n", customer.name)
	if name == "tom" && password == "12345" {
		return true
	} else {
		return false
	}
}

func main() {

	cus := Customer{
		name: "tom",
	}
	b := cus.login("tom", "12345")
	fmt.Printf("%v\n", b)

	/* per := Person{
		name : "tom",
	}
	per.eat()
	per.sleep() */
}
