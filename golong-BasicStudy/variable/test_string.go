package main

import (
	"fmt"
	"strings"
	// "bytes"
)

func main() {

	/* var s string = "hello world"
	var s1 = "Hello world"
	s2 := "Hello World"

	fmt.Printf("s: %v\n", s)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)

	s4 := `
	line 1
	line 2
	line 3
	`
	fmt.Printf("s4: %v\n", s4) */

	//字符串连接
	/* s1 := "tom"
	s2 := "20"
	msg := s1 + s2
	fmt.Printf("msg: %v\n", msg) */

	//字符串连接
	/* s1 := "tom"
	s2 := "20"
	msg := fmt.Sprintf("name = %s, age = %s", s1, s2)
	fmt.Printf("msg: %v\n", msg) */

	/* name := "tom"
	age := "20"
	s := strings.Join([]string{name,age},",")
	fmt.Printf("s: %v\n", s) */

	/* var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer.string(): %v\n",buffer.String()) */

	//转义字符
	/* s := "hello"
	print(s + "\n")
	print(s) */

	//字符串的切片操作
	/* s := "hello world"
	fmt.Printf("s[2]: %c\n", s[2])
	fmt.Printf("s[2:5]: %v\n", s[2:5])
	fmt.Printf("s[2:]: %v\n", s[2:])
	fmt.Printf("s[:5]: %v\n", s[:5]) */

	//字符串函数
	s := "Hello World"
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, " "))
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))
	fmt.Printf("strings.ToLower(s): %v\n", strings.ToLower(s))
	fmt.Printf("strings.ToUpper(s): %v\n", strings.ToUpper(s))
	fmt.Printf("strings.HasPrefix(\"Hello\"): %v\n", strings.HasPrefix(s, "Hello"))
	fmt.Printf("strings.HasSuffix(\"world\"): %v\n", strings.HasSuffix(s, "world"))
	fmt.Printf("strings.Index(s, \"l\"): %v\n", strings.Index(s, "l"))
	fmt.Printf("strings.LastIndex(s, \"l\"): %v\n", strings.LastIndex(s, "l"))

}
