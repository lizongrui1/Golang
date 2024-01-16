package main

import (
	"StudentManagementSystem/module"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := module.InitDB()
	if err != nil {
		log.Fatalf("数据库初始化失败: %v\n", err)
		return
	}
	defer db.Close()

	// 设置静态文件服务
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 设置路由
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/query", module.QueryRowHandler)
	http.HandleFunc("/insert", module.InsertRowHandler)
	http.HandleFunc("/update", module.UpdateRowHandler)
	http.HandleFunc("/delete", module.DeleteRowHandler)

	// 启动 HTTP 服务器
	fmt.Println("服务器运行在 http://localhost:8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// homeHandler 用于处理主页请求并显示主页 HTML
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 这里填充用于生成主页的代码，例如加载 HTML 文件
	http.ServeFile(w, r, "./templates/index.html")
}
