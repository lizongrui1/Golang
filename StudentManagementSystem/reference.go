package main

//
//import (
//	"database/sql"//用于操作SQL数据库
//	"fmt"//用于格式化字符串输出，
//	"log"//用于输出日志
//	"net/http"//用于处理HTTP请求和响应
//	"strconv"//用于字符串和其他数据类型的转换，
//	"text/template"//用于HTML模板解析；
//
//	_ "github.com/go-sql-driver/mysql"
//	//使用空标识符导入mysql驱动，表示在引入该包时，
//	//仅执行该包的init函数，注册相应的数据库驱动。
//)
//
//// 数据库连接信息
////使用const关键字定义常量，表示数据库连接的相关信息，
////如数据库驱动、用户名、密码、主机地址、端口号和数据库名等
//const (
//	dbDriver   = "mysql"
//	dbUser     = "root"
//	dbPassword = "password123"
//	dbHost     = "127.0.0.1"
//	dbPort     = 3306
//	dbName     = "students"
//)
//
//// 学生信息结构体
////使用type关键字定义一个Student结构体类型，表示学生的信息，包括学生ID、姓名和分数等信息。
//type Student struct {
//	ID     int
//	Name   string
//	Score  int
//}
///*使用var关键字定义变量tpl和db，分别用于保存HTML模板和数据库连接。
//在init()函数中，使用template.Must()方法解析templates目录下的所有HTML文件，然后使用sql.Open()函数打开数据库连接。
//fmt.Sprintf()方法将数据库连接信息格式化为一个字符串，然后传递给sql.Open()函数实现连接数据库。
//如果连接失败，使用log.Fatalf()方法输出错误信息并终止程序；如果连接成功，使用log.Println()方法输出日志信息表示连接已建立。*/
//var tpl *template.Template
//var db *sql.DB
//
//func init() {
//	// 使用Template包解析对应的HTML文件
//	tpl = template.Must(template.ParseGlob("templates/*.html"))
//
//	// 连接数据库
//	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
//	var err error
//	db, err = sql.Open(dbDriver, dataSourceName)
//	if err != nil {
//		log.Fatalf("Failed to connect to database: %v", err)
//	}
//	log.Println("Database connection established.")
//}
//
//// 获取所有学生信息的Handler
///*getAllStudents()方法使用db.Query()方法执行SQL查询语句，获取所有学生记录的数据。
//使用defer rows.Close()方法，确保在函数退出时关闭连接。
//创建一个students空切片，使用rows.Next()方法遍历每条学生记录，并使用rows.Scan()方法将学生信息分别存储在Student结构体对象中。
//使用append()方法将每个学生对象添加到students切片中。
//最后，使用rows.Err()方法检查是否有错误发生，并返回学生数据和错误信息。*/
//func allStudentsHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "POST" {
//		r.ParseForm()
//		name := r.FormValue("name")
//		score, err := strconv.Atoi(r.FormValue("score"))
//		if err != nil {
//			http.Error(w, "Invalid score value", http.StatusBadRequest)
//		}
//		addStudent(name, score)
//	}
//
//	students, err := getAllStudents()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	tpl.ExecuteTemplate(w, "index.html", students)
//}
//
//// 添加学生信息的Handler
//func addStudentHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "POST" {
//		r.ParseForm()
//		name := r.FormValue("name")
//		score, err := strconv.Atoi(r.FormValue("score"))
//		if err != nil {
//			http.Error(w, "Invalid score value", http.StatusBadRequest)
//		}
//		addStudent(name, score)
//		http.Redirect(w, r, "/", http.StatusSeeOther)
//	} else {
//		tpl.ExecuteTemplate(w, "add.html", nil)
//	}
//}
//
//// 修改学生信息的Handler
//func editStudentHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "POST" {
//		r.ParseForm()
//		id, err := strconv.Atoi(r.FormValue("id"))
//		if err != nil {
//			http.Error(w, "Invalid student ID value", http.StatusBadRequest)
//			return
//		}
//		name := r.FormValue("name")
//		score, err := strconv.Atoi(r.FormValue("score"))
//		if err != nil {
//			http.Error(w, "Invalid score value", http.StatusBadRequest)
//			return
//		}
//		editStudent(id, name, score)
//		http.Redirect(w, r, "/", http.StatusSeeOther)
//	} else {
//		studentID := r.FormValue("id")
//		if studentID == "" {
//			http.Error(w, "Student ID not provided", http.StatusBadRequest)
//			return
//		}
//		id, err := strconv.Atoi(studentID)
//		if err != nil {
//			http.Error(w, "Invalid student ID value", http.StatusBadRequest)
//			return
//		}
//		student, err := getStudentByID(id)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		tpl.ExecuteTemplate(w, "update.html", student)
//	}
//}
//
//// 删除学生信息的Handler
//func deleteStudentHandler(w http.ResponseWriter, r *http.Request) {
//	id, err := strconv.Atoi(r.URL.Query().Get(“id”))
//	if err != nil {
//		http.NotFound(w, r)
//		return
//	}
//	if err := deleteStudent(id); err != nil {
//		http.NotFound(w, r)
//		return
//	}
//	http.Redirect(w, r, “/”, http.StatusSeeOther)
//}
//
//func main() {
//	defer db.Close()
//	http.HandleFunc("/", indexHandler)
//	http.HandleFunc("/add", addHandler)
//	http.HandleFunc("/edit", editHandler)
//	http.HandleFunc("/delete", deleteHandler)
//
//	log.Println("Starting server...")
//	http.ListenAndServe(":8080", nil)
//}
