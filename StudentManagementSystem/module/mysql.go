package module

import (
	"database/sql"
	"html/template"
)

var tpl *template.Template

func InitDB() (*sql.DB, error) {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/studb")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	_, err = db.Exec("USE studb")
	if err != nil {
		return nil, err
	}
	return db, nil
}
