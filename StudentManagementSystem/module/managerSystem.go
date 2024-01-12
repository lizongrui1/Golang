package module

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

var db = db_init()

type myUsualType interface{}

/*func InitDB() error {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("数据库连接成功...")

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}*/

//create table users(
//id int primary key auto_increment,
//username varchar(100) not null unique,
//password varchar(100) not null,
//email varchar(100)
//)

func ShowMenu() {
	fmt.Println("-----------------------------")
	fmt.Println("欢迎使用学生管理系统 V1.0")
	fmt.Println("1.显示所有学生的信息")
	fmt.Println("2.显示指定学生的信息")
	fmt.Println("3.学生信息修改")
	fmt.Println("4.新增学生信息")
	fmt.Println("5.删除学生信息")
	fmt.Println()
	fmt.Println("0.退出系统")
	fmt.Println("-----------------------------")
}

func FunctionChoose(button int) error {
	var err error
	switch button {
	case 1:
		// 查询所有学生信息
		err = queryMultiRow()
		//if err != nil {
		//	return err
		//}
		checkError(err) //后面if err..未进行修改

	case 2:
		// 查询单个学生信息
		var id int
		fmt.Printf("请输入查询的学生学号：")
		_, err = fmt.Scan(&id)
		if err != nil {
			return err
		}
		err = queryRow(id)
		if err != nil {
			return err
		}

	case 3:
		var (
			id       int
			item     string
			newValue string //获取的newValue初始都设置成string类型，得到值后，进行类型转换
			a        myUsualType
		)
		fmt.Printf("请输入修改信息的学生学号：")
		_, err = fmt.Scan(&id)
		if err != nil {
			return err
		}
		//如果学生信息不存在，需要报错返回 查询学生信息的操作
		err = queryRow(id)
		if err != nil {
			return err
		}
		fmt.Printf("请输入修改的信息：（字段 新值）")
		_, err = fmt.Scan(&item, &newValue)
		if err != nil {
			return err
		}
		//将获取的item字段更改为小写
		item = strings.ToLower(item)
		//newValue的类型转换
		if item == "id" || item == "age" || item == "chinese" || item == "math" || item == "english" {
			a, err = strconv.ParseInt(newValue, 10, 0) //string 转int64
			if err != nil {
				return err
			}
		} else {
			a = fmt.Sprintf("'%v'", newValue)
		}
		err = updateRow(id, item, a)
		if err != nil {
			return err
		}

	case 4:
		var s Student
		fmt.Println("请输入新增学生信息：\n", "(学号，姓名，成绩)")
		_, err = fmt.Scan(&s.Id, &s.Name, &s.Score)
		if err != nil {
			return err
		}
		err = insertRow(s.Id, s.Name, s.Score)
		if err != nil {
			return err
		}

	case 5:
		// 删除学生信息
		var id int
		fmt.Printf("请输入删除的学生学号：")
		_, err = fmt.Scan(&id)
		if err != nil {
			return err
		}
		err = deleteRow(id)
		if err != nil {
			return err
		}
	default:
		err = errors.New("输入错误！")
		return err
	}
	return err
}

// 功能。。。
func queryRow(id int) error {
	var stu Student // 假设 Student 是一个定义好的结构体
	err := db.QueryRow("select * from `users` where id=?", id).Scan(&stu.Id, &stu.Name, &stu.Score)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return err
	}
	fmt.Println("查询成功!")
	fmt.Printf("学号: %d, 姓名: %s, 分数: %s\n", stu.Id, stu.Name, stu.Score)
	return nil
}

func queryMultiRow() error {
	rows, err := db.Query("select * from users")
	if err != nil {
		fmt.Println("query failed, err:%v\n", err)
		return nil
	}
	defer rows.Close()

	fmt.Println("查询成功!")
	for rows.Next() {
		err := rows.Scan(&stu.Id, &stu.Name, &stu.Score)
		if err != nil {
			fmt.Println("scan failed, err:%v\n", err)
			return nil
		}
		fmt.Printf("学号: %d, 姓名: %s, 分数: %s\n", stu.Id, stu.Name, stu.Score)
	}
	return nil
}

func insertRow(id int, name string, score int) error {
	ret, err := db.Exec("insert into users(id, name, score) values (?, ?, ?)", id, name, score)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	insertedId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return err
	}
	fmt.Printf("加入成功, 学号为 %d\n", insertedId)
	return nil
}

func updateRow(id int, item string, newValue myUsualType) error {
	sqlStr := fmt.Sprintf("update users set%s = ? where id = ?", item, id)
	ret, err := db.Exec(sqlStr, newValue, id)
	if err != nil {
		fmt.Printf("更新失败, errr:%v\n", err)
		return nil
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return nil
	}
	fmt.Printf("更新成功, affected rows: %d\n", n)
	return nil
}

func deleteRow(id int) (err error) {
	ret, err := db.Exec("delete from users where id = ?", id)
	if err != nil {
		fmt.Printf("删除失败, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("删除成功, affected rows:%d\n", n)
	return err
}
