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

var stu = new(Student)

<<<<<<< HEAD
// 查看学生
func queryRow(id int) error {
	var stu Student // 假设 Student 是一个定义好的结构体
	err := db.QueryRow("select * from `sms` where sms.id=?", id).Scan(&stu.Id, &stu.Name, &stu.Score)
=======
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

//func FunctionChoose(button int) error {
//	var err error
//	switch button {
//	case 1:
//		// 查询所有学生信息
//		students, err := queryMultiRow()
//		if err != nil {
//			return err
//		}
//		for _, student := range students {
//			fmt.Printf("学号: %d, 姓名: %s, 分数: %d\n", student.Id, student.Name, student.Score)
//		}
//
//	case 2:
//		// 查询指定学生信息
//		var id int
//		fmt.Printf("请输入查询的学生学号：")
//		_, err = fmt.Scan(&id)
//		if err != nil {
//			return err
//		}
//		err = queryRow(id)
//		if err != nil {
//			return err
//		}
//
//	case 3:
//		var (
//			id       int
//			item     string
//			newValue string //获取的newValue初始都设置成string类型，得到值后，进行类型转换
//			a        myUsualType
//		)
//		fmt.Printf("请输入修改信息的学生学号：")
//		_, err = fmt.Scan(&id)
//		if err != nil {
//			return err
//		}
//		//如果学生信息不存在，需要报错返回 查询学生信息的操作
//		err = queryRow(id)
//		if err != nil {
//			return err
//		}
//		fmt.Printf("请输入修改的信息：（字段 新值）")
//		_, err = fmt.Scan(&item, &newValue)
//		if err != nil {
//			return err
//		}
//		//将获取的item字段更改为小写
//		item = strings.ToLower(item)
//		//newValue的类型转换
//		if item == "id" || item == "age" || item == "chinese" || item == "math" || item == "english" {
//			a, err = strconv.ParseInt(newValue, 10, 0) //string 转int64
//			if err != nil {
//				return err
//			}
//		} else {
//			a = fmt.Sprintf("'%v'", newValue)
//		}
//		err = updateRow(id, item, a)
//		if err != nil {
//			return err
//		}
//
//	case 4:
//		var s Student
//		fmt.Println("请输入新增学生信息：\n", "(学号，姓名，成绩)")
//		_, err = fmt.Scan(&s.Id, &s.Name, &s.Score)
//		if err != nil {
//			return err
//		}
//		err = insertRow(s.Name, s.Score)
//		if err != nil {
//			return err
//		}
//
//	case 5:
//		// 删除学生信息
//		var id int
//		fmt.Printf("请输入删除的学生学号：")
//		_, err = fmt.Scan(&id)
//		if err != nil {
//			return err
//		}
//		err = deleteRow(id)
//		if err != nil {
//			return err
//		}
//	default:
//		err = errors.New("输入错误！")
//		return err
//	}
//	return err
//}

// 查看学生
func queryRow(id int) error {
	var stu Student // 假设 Student 是一个定义好的结构体
	err := db.QueryRow("select * from `sms` where sms.sid=?", id).Scan(&stu.Id, &stu.Name, &stu.Score)
>>>>>>> 6f45ca077e45fe2b265c28107e07681450dc02a6
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return err
	}
	//defer db.Close()
	fmt.Println("查询成功!")
	fmt.Printf("学号: %d, 姓名: %s, 分数: %d\n", stu.Id, stu.Name, stu.Score)
	return nil
}

// 多行查看
func queryMultiRow() ([]Student, error) {
<<<<<<< HEAD
	rows, err := db.Query("SELECT id, name, score FROM sms")
=======
	rows, err := db.Query("SELECT sid, sname, score FROM sms")
>>>>>>> 6f45ca077e45fe2b265c28107e07681450dc02a6
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
	ret, err := db.Exec("insert into sms(sid, sname, score) values (?, ?, ?)", name, score)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	insertedId, err := ret.LastInsertId() //获取插入操作的最后插入的自增主键
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return err
	}
	fmt.Printf("加入成功, 学号为 %d\n", insertedId)
	return nil
}

// 修改学生
func updateRow(id int, newValue myUsualType) error {
	sqlStr := "UPDATE sms SET score = ? WHERE id = ?"
	ret, err := db.Exec(sqlStr, newValue, id)
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
	fmt.Printf("更新成功")
	return nil
}

// 删除学生
func deleteRow(id int) (err error) {
	ret, err := db.Exec("delete from sms where sid = ?", id)
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
