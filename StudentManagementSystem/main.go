package main

import (
	"StudentManagementSystem/module"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var input int
	db, err := module.InitDB()
	if err != nil {
		log.Fatalf("数据库初始化失败: %v\n", err)
		return
	}
	defer db.Close()

	for {
		module.ShowMenu()
		fmt.Println("请输入选项：")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("输入错误: %v\n", err)
			continue // 输入错误时，继续循环
		}

		if input == 0 {
			fmt.Println("退出系统...")
			break // 退出循环
		}

		err = module.FunctionChoose(input) // 调用FunctionChoose处理用户输入
		if err != nil {
			fmt.Printf("操作错误: %v\n", err) // 处理FunctionChoose返回的错误
		}
	}

	// 将 queryRowHandler 函数绑定到特定路由
	http.HandleFunc("/query", module.QueryRowHandler)
	http.HandleFunc("/insert", module.InsertRowHandler)
	http.HandleFunc("/update", module.UpdateRowHandler)
	http.HandleFunc("/delete", module.DeleteRowHandler)

	// 启动 HTTP 服务器
	fmt.Println("服务器运行在 http://localhost:8000")
	http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
