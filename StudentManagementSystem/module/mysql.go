package module

import (
	"database/sql"
)

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
}
