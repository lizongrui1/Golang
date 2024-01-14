package main

import (
	"StudentManagementSystem/module"
	"fmt"
	"net/http"
)

func main() {
	var input int
	err := module.InitDB()
	if err != nil {
		fmt.Printf("Failed to initialize DB: %v\n", err)
		return
	}

	for {
		module.ShowMenu()
		fmt.Println("请输入选项：")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("输入错误: %v\n", err)
			continue
		}

		choice := input
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

	// 将 queryRowHandler 函数绑定到特定路由
	http.HandleFunc("/query", module.QueryRowHandler)

	// 启动 HTTP 服务器
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
