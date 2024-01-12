package main

import (
	"StudentManagementSystem/module"
	"fmt"
	"strconv"
)

func main() {
	err := module.InitDB()
	if err != nil {
		fmt.Printf("Failed to initialize DB: %v\n", err)
		return
	}

	//student := module.NewStudent(6, "xiaoli", 10086)
	//fmt.Printf("Student ID: %d, Name: %s, Score: %d\n", student.Id, student.Name, student.Score)

	for {
		module.ShowMenu()
		fmt.Println("请输入选项：")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("输入错误: %v\n", err)
			continue
		}

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("请输入有效的数字")
			continue
		}

		if choice == 0 {
			fmt.Println("退出系统...")
			break
		}

		err = module.FunctionChoose(choice)
		if err != nil {
			fmt.Printf("操作错误: %v\n", err)
		}
	}
}
