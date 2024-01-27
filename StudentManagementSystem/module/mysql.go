package module

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
)

<<<<<<< HEAD
var tpl *template.Template

func InitDB() (*sql.DB, error) {
	tpl = template.Must(template.ParseGlob("D:/goProject/gotogit/StudentManagementSystem/module/templates/*.html"))
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/studb?charset=utf8")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	_, err = db.Exec("USE studb")
	if err != nil {
		return nil, err
	}
	return db, nil
=======
func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	//if err != nil {
	//	panic(err)
	//}
	checkError(err)
	err = db.Ping()
	checkError(err)
	db.Exec("use test")
	defer db.Close()
	return db
>>>>>>> 688cf5ed53b85b39b20b54793577ffe795686929
}
