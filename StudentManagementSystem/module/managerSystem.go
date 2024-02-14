package module

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var db, _ = InitDB()

var rdb *redis.Client
var ctx = context.Background()

type myUsualType interface{}

type Student struct {
	Number int `json:"number"`
	Name   string
	Score  int
}

func register(number string, password string) (err error) {
	currentTime := time.Now()
	ret, err := db.Exec("INSERT INTO stu (student_id, password) VALUES (?, ?)", number, password)
	if err != nil {
		log.Printf("学生账号添加失败: %v\n", err)
		return
	}
	newID, err := ret.LastInsertId()
	if err != nil {
		log.Printf("新注册学生ID失败: %v\n", err)
	}
	log.Printf("%s注册成功, 新注册的学生学号为：%d\n", currentTime.Format("2006/01/02 15:04:05"), newID)
	return
}

// 查看学生
//
//	func queryRow(number int) (student Student, err error) {
//		var stu Student
//		err = db.QueryRow("SELECT number, name, score FROM sms WHERE number = ?", number).Scan(&stu.Number, &stu.Name, &stu.Score)
//		if err != nil {
//			fmt.Printf("查询失败, err: %v\n", err)
//			return
//		}
//		return stu, nil
//	}
func queryRow(number int) (student Student, err error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	val, err := rdb.Get(timeoutCtx, fmt.Sprintf("student:%d", number)).Result()
	if err == redis.Nil {
		err = db.QueryRow("SELECT number, name, score FROM sms WHERE number = ?", number).Scan(&student.Number, &student.Name, &student.Score)
		if err != nil {
			log.Fatalf("查询失败, err: %v\n", err)
			return
		}
		studentJSON, _ := json.Marshal(student)
		err = rdb.Set(timeoutCtx, fmt.Sprintf("student:%d", number), studentJSON, 30*time.Minute).Err()
		if err != nil {
			log.Printf("缓存设置失败, err: %v\n", err)
		}
		return student, nil
	} else if err != nil {
		log.Printf("从Redis查询失败, err: %v\n", err)
		return
	} else {
		err = json.Unmarshal([]byte(val), &student)
		if err != nil {
			log.Printf("反序列化失败, err: %v\n", err)
			return
		}
		return student, nil
	}
}

// 全部查看
func queryMultiRow() ([]Student, error) {
	var students []Student
	ret, err := db.Query("SELECT number, name, score FROM sms")
	if err != nil {
		log.Printf("查询失败, err:%v\n", err)
		return nil, err
	}
	defer ret.Close()
	for ret.Next() {
		var stu Student
		err := ret.Scan(&stu.Number, &stu.Name, &stu.Score)
		if err != nil {
			log.Printf("赋值失败, err:%v\n", err)
			continue
		}
		students = append(students, stu)
	}
	if err := ret.Err(); err != nil {
		log.Printf("迭代失败, err:%v\n", err)
		return nil, err
	}
	return students, nil
}

// 增加学生
func insertRow(number int, name string, score int) (err error) {
	currentTime := time.Now()
	ret, err := db.Exec("INSERT INTO sms (number, name, score) VALUES (?, ?, ?)", number, name, score)
	if err != nil {
		fmt.Printf("添加失败, err:%v\n", err)
		return
	}
	insertedId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("获取插入ID失败, err:%v\n", err)
		return
	}
	log.Printf("%s 加入成功, 新加入的学生序号为：%d\n", currentTime.Format("2006/01/02 15:04:05"), insertedId)
	return
}

// 修改学生
func updateRow(number int, newScore myUsualType) (err error) {
	sqlStr := "UPDATE sms SET score = ? WHERE number = ?"
	ret, err := db.Exec(sqlStr, newScore, number)
	if err != nil {
		log.Fatalf("更新失败, error: %v\n", err)
		return
	}
	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		log.Printf("获取更新行数时发生错误: %v\n", err)
		return
	}
	if rowsAffected == 0 {
		fmt.Println("没有找到对应的学号, 未进行更新")
		return
	}
	fmt.Printf("更新成功, 受影响行数:%d\n", rowsAffected)
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	student, err := queryRow(number)
	if err != nil {
		log.Printf("从数据库获取更新后的学生信息失败, err: %v\n", err)
		return
	}

	studentJSON, err := json.Marshal(student)
	if err != nil {
		log.Printf("序列化学生信息失败, err: %v\n", err)
		return
	}

	err = rdb.Set(timeoutCtx, fmt.Sprintf("student:%d", number), studentJSON, 30*time.Minute).Err()
	if err != nil {
		log.Printf("更新Redis缓存失败, err: %v\n", err)
		return
	}
	return
}

// 删除学生
func deleteRow(number int) (err error) {
	currentTime := time.Now()
	// 首先检查学生是否存在
	_, err = queryRow(number)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("没有找到学号为 %d 的学生", number)
		}
		fmt.Printf("查询学生时出错, err: %v\n", err)
		return
	}
	ret, err := db.Exec("DELETE FROM sms WHERE number = ?", number)
	if err != nil {
		fmt.Printf("删除失败, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	if n == 0 {
		return fmt.Errorf("没有学号为 %d 的学生", number)
	}
	fmt.Printf("%s 删除成功, 删除的学生学号为：%d", currentTime.Format("2006/01/02 15:04:05"), number)
	return
}
