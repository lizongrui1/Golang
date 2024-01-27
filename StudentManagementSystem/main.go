package main

import (
	"StudentManagementSystem/module"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
<<<<<<< HEAD
	db, err := module.InitDB()
=======
	err := module.InitDB()
>>>>>>> 688cf5ed53b85b39b20b54793577ffe795686929
	if err != nil {
		log.Fatalf("数据库初始化失败: %v\n", err)
		return
	}
	defer db.Close()

	// 设置静态文件服务
	fs := http.FileServer(http.Dir("./module/templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

<<<<<<< HEAD
	http.HandleFunc("/", module.HomeHandler)
	http.HandleFunc("/query", module.QueryRowHandler)
	http.HandleFunc("/insert", module.InsertRowHandler)
	http.HandleFunc("/update", module.UpdateRowHandler)
	http.HandleFunc("/delete", module.DeleteRowHandler)

	fmt.Println("服务器运行在 http://localhost:8080")
	err = http.ListenAndServe("localhost:8080", nil)
=======
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

	// 将 queryRowHandler 函数绑定到特定路由
	http.HandleFunc("/query", QueryRowHandler)

	// 启动 HTTP 服务器
	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
>>>>>>> 688cf5ed53b85b39b20b54793577ffe795686929
	if err != nil {
		log.Fatal("发生错误:", err)
	}
}
