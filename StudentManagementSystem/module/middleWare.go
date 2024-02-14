package module

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

//var tpl *template.Template

func InitDB() (*sql.DB, error) {
	//tpl = template.Must(template.ParseGlob("./module/templates/*.html"))
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/studb?charset=utf8")
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
	return db, nil
}

func InitRDB() error {
	rdb = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("无法连接到Redis: %v", err)
	}

	var err error
	db, err = InitDB()
	if err != nil {
		log.Fatalf("无法连接到MySQL数据库: %v", err)
	}
	return nil
}
