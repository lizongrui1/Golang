package module

import (
	"database/sql"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/studb")
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
