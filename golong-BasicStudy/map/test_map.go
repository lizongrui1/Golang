package main

import "fmt"

/*func test1() {
	var m1 map[string]string //类型声明
	m1 = make(map[string]string)
	fmt.Printf("%v\n", m1)
	fmt.Printf("%T\n", m1)
}*/

/*func test2() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "gmail.com"}
	fmt.Printf("%v\n", m1)

	m2 := make(map[string]string)
	m2["name"] = "Sam"
	m2["age"] = "23"
	m2["email"] = "qq.com"
	fmt.Printf("%v\n", m2)
}*/

/*func test3() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "gmail.com"}
	var key = "name"
	var value = m1[key]
	fmt.Printf("%v\n", value)
}*/

func test4() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "gmail.com"}
	var k1 = "name"
	var k2 = "age1"
	v, ok := m1[k1]
	fmt.Printf("%v\n", v)  //tom
	fmt.Printf("%v\n", ok) //true
	fmt.Println("------------")
	v, ok = m1[k2]
	fmt.Printf("%v\n", v)  //空
	fmt.Printf("%v\n", ok) //false
}

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}
