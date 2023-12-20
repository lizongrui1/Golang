package main

import "fmt"

/*func f4() {
var a = [...]int{1, 2, 3}
for i, v := range a {
	/*fmt.Printf("%v\n", i) //数组索引
	fmt.Printf("%v\n", v) //值
	fmt.Printf("%v：%v ", i, v)
	}
	}*/

/*func f4() {
	var s = []int{1, 2, 3} //[]里面什么都不加，是切片
	for _, v := range s {
		fmt.Println("v：", v)
	}
}*/

/*func f4() {
	//  key：value
	m := make(map[string]string, 0)
	m["name"] = "lzr"
	m["age"] = "20"
	m["email"] = "lizongrui99@gmail.com"
	for key, value := range m {
		fmt.Println(key, value)
	}
}*/

func f4() {
	s := "hello"
	for _, v := range s {
		fmt.Println(string(v))
	}
}

func main() {
	f4()
}
