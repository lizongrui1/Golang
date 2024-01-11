package module

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func InitDB() (err error) {
	db, err = sql.Open("test", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v\n", err)
		}
	}(db)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

//query := `
//   CREATE TABLE users (
//       id INT AUTO_INCREMENT,
//       username TEXT NOT NULL,
//       password TEXT NOT NULL,
//       created_at DATETIME,
//       PRIMARY KEY (id)
//   );`
//// 执行后一定要检查err
//_, err := db.Exec(query)

func showMenu() {
	fmt.Println("******************************")
	fmt.Println("欢迎使用[学生管理系统] V1.0")
	fmt.Println("1.添加学生")
	fmt.Println("2.显示全部")
	fmt.Println("3.查询学生")
	fmt.Println("4.修改学生")
	fmt.Println("5.删除学生")
	fmt.Println()
	fmt.Println("0.退出系统")
	fmt.Println("******************************")
}

func queryRow() {
	err := db.QueryRow("select id, name, phone from `users` where id=?", 1).Scan(&stu.Id, &stu.Name, &stu.Score)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	fmt.Println("Query success!")
	fmt.Printf("id: %d, name: %s, phone: %s\n", stu.Id, stu.Name, stu.Score)
}

func queryMultiRow() {
	rows, err := db.Query("select * from users")
	if err != nil {
		fmt.Println("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	fmt.Println("query success!")
	for rows.Next() {
		err := rows.Scan(&stu.Id, &stu.Name, &stu.Score)
		if err != nil {
			fmt.Println("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id: %d, name: %s, phone: %s\n", stu.Id, stu.Name, stu.Score)
	}
}

func insterRow() {
	ret, err := db.Exec("insert into users(name, phone) values (?,?)", "zhaoliu", "44444")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d\n", id)
}

func updateRow() {
	ret, err := db.Exec("update users set name = ? where id = ?", "haha", 4)
	if err != nil {
		fmt.Printf("update failed, errr:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)
}

func deleteRow() {
	ret, err := db.Exec("delete from users where id = ?", 2)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
