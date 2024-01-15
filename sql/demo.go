package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB // 定义全局变量用来保存数据库连接对象，不需要初始化为 new(sql.DB)
var db *sql.DB // 不需要初始化为 new(sql.DB)

func initDB() (err error) {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	Id    int
	Name  string
	Phone string
}

var u *User = new(User)

func queryRow() {
	err := db.QueryRow("select id, name, phone from `users` where id=?", 1).Scan(&u.Id, &u.Name, &u.Phone)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	fmt.Println("Query success!")
	fmt.Printf("id: %d, name: %s, phone: %s\n", u.Id, u.Name, u.Phone)
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
		err := rows.Scan(&u.Id, &u.Name, &u.Phone)
		if err != nil {
			fmt.Println("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id: %d, name: %s, phone: %s\n", u.Id, u.Name, u.Phone)
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
	ret, err := db.Exec("update users set name, = ? where id = ?", "haha", 4)
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

func main() {
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/*	queryRow()
		fmt.Println("===============")*/
	queryMultiRow()
	/*fmt.Println("===============")
	insterRow()*/
	/*fmt.Println("===============")
	updateRow()*/
	fmt.Println("===============")
	//deleteRow()
}
