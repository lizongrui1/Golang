package module

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
)

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
}
