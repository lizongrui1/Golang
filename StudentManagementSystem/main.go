package main

import (
	"StudentManagementSystem/module"
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"os"
)

func main() {
	//mysql
	db, err := module.InitDB()
	if err != nil {
		log.Fatalf("数据库初始化失败，err:%v\n", err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("数据库关闭失败，err:%v\n", err)
		}
	}(db)
	//redis
	err = module.InitRDB()
	if err := module.InitRDB(); err != nil {
		log.Fatalf("初始化失败: %v", err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// 启用跟踪仪器
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		log.Fatalf("无法为Redis启用跟踪: %v", err)
	}
	// 启用指标仪器
	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		log.Fatalf("无法为Redis启用指标: %v", err)
	}

	// 创建日志文件
	file, err := os.OpenFile("student_management_system.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) // 如果文件不存在就创建它，然后以只写模式打开，且写入的数据追加到文件末尾
	if err != nil {
		log.Fatal(err)
	}
	//logger = log.New(file, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("日志关闭失败，err: %v\n", err)
		}
	}(file)

	// 设置日志输出到文件
	log.SetOutput(file)

	// 设置静态文件服务
	fs := http.FileServer(http.Dir("./module/templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/login", module.LoginHandler)
	http.HandleFunc("/home", module.HomeHandler)
	http.HandleFunc("/query", module.QueryRowHandler)
	http.HandleFunc("/queryAll", module.QueryAllRowHandler)
	http.HandleFunc("/insert", module.InsertRowHandler)
	http.HandleFunc("/update", module.UpdateRowHandler)
	http.HandleFunc("/delete", module.DeleteRowHandler)
	http.HandleFunc("/register", module.RegisterStudentHandler)
	http.HandleFunc("/studentPage", module.StudentPageHandler)
	http.HandleFunc("/studentSelect", func(w http.ResponseWriter, r *http.Request) {
		module.StudentSelectHandler(w, r, db, rdb)
	})

	fmt.Println("学生管理系统运行在： http://127.0.0.1:8080， 按 CTRL + C 退出系统。")
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("发生错误:", err)
	}
}
