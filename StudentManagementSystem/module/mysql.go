package module

import (
	"database/sql"
)

func db_init() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	//if err != nil {
	//	panic(err)
	//}
	checkError(err)

	err = db.Ping()
	checkError(err)

	query := `
	CREATE TABLE users (
	   id INT AUTO_INCREMENT PRIMARY KEY,
	   name TEXT NOT NULL,
	   score INT AUTO_INCREMENT,
	   PRIMARY KEY (id)
	);`
	_, err = db.Exec(query)
	checkError(err)

	return db
}
