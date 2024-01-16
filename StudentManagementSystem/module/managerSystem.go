package module

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db, _ = InitDB()

type myUsualType interface{}

type Student struct {
	Id    int
	Name  string
	Score int
}

//var stu = new(Student)

// 查看学生
func queryRow(id int) (student Student, err error) {
	var stu Student
	err = db.QueryRow("select * from `sms` where id=?", id).Scan(&stu.Id, &stu.Name, &stu.Score)
	if err != nil {
		fmt.Printf("查询失败, err: %v\n", err)
		return
	}
	fmt.Println("查询成功!")
	fmt.Printf("学号: %d, 姓名: %s, 分数: %d\n", stu.Id, stu.Name, stu.Score)
	return stu, nil
}

// 全部查看
func queryMultiRow() ([]Student, error) {
	rows, err := db.Query("SELECT id, name, score FROM sms")
	if err != nil {
		fmt.Printf("查询失败, err:%v\n", err)
		return nil, err
	}
	defer rows.Close()

	var students []Student // 创建一个空的Student切片用于存储学生数据

	for rows.Next() {
		var stu Student // 使用Student结构体来存储每行的数据
		err := rows.Scan(&stu.Id, &stu.Name, &stu.Score)
		if err != nil {
			fmt.Printf("赋值失败, err:%v\n", err)
			continue // 发生错误时继续处理下一行
		}
		students = append(students, stu) // 将学生对象添加到切片中
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("iteration failed, err:%v\n", err)
		return nil, err
	}

	return students, nil // 返回学生切片和错误信息
}

// 增加学生
func insertRow(name string, score int) error {
	ret, err := db.Exec("INSERT INTO sms (name, score) VALUES (?, ?)", name, score)
	if err != nil {
		fmt.Printf("添加失败, err:%v\n", err)
		return err
	}
	insertedId, err := ret.LastInsertId() //获取插入操作的最后插入的自增主键
	if err != nil {
		fmt.Printf("获取最后插入ID失败, err:%v\n", err)
		return err
	}
	fmt.Printf("加入成功, 新加入的学生编号为：%d\n", insertedId)
	return nil
}

// 修改学生
func updateRow(name string, newValue myUsualType) error {
	sqlStr := "UPDATE sms SET score = ? WHERE name = ?"
	ret, err := db.Exec(sqlStr, newValue, name)
	if err != nil {
		fmt.Printf("更新失败, error: %v\n", err)
		return err
	}
	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("获取更新行数时发生错误: %v\n", err)
		return err
	}
	if rowsAffected == 0 {
		fmt.Println("没有找到对应的ID, 未进行更新")
		return nil
	}
	fmt.Printf("更新成功, 受影响行数:%d\n", rowsAffected)
	return nil
}

// 删除学生
func deleteRow(id int) (err error) {
	ret, err := db.Exec("delete from sms where id = ?", id)
	if err != nil {
		fmt.Printf("删除失败, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("删除成功, 受影响行数:%d\n", n)
	return err
}
