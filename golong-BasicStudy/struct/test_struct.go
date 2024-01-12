package main

//结构体
import "fmt"

type Person struct {
	id    int
	name  string
	age   int
	email string
}

type Customer struct {
	id, age     int
	name, email string
}

func main() {
	/* var tom Person //var i int
	fmt.Printf("tom: %v\n", tom)

	var kit Customer
	fmt.Printf("kit: %v\n", kit) */

	/* var tom Person
	tom.id = 10086
	tom.name = "tom"
	tom.age = 18
	tom.email = "12345@gamil.com"
	fmt.Printf("tom: %v\n", tom)

	fmt.Printf("tom.age: %v\n", tom.age) */

	//匿名结构体
	var tom struct {
		id   int
		name string
		age  int
	}

	tom.id = 10086
	tom.name = "tom"
	tom.age = 18
	fmt.Printf("tom: %v\n", tom)

}
