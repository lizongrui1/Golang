package module

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

//type DataStore interface {
//	QueryRow(number int) (Student, error)
//}
//
//type MockDataStore struct {
//}

func initDB() *sql.DB {
	//wd, err := os.Getwd()
	//if err != nil {
	//	log.Fatalf("无法获取当前工作目录: %s", err)
	//}
	//fmt.Println("当前工作目录:", wd)
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/studb?charset=utf8")
	if err != nil {
		log.Fatalf("无法连接数据库: %s", err)
	}

	return db
}

func setup(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO sms (number, name, score) VALUES (9997, '测试1', 88), (9998, '测试2', 92), (9999, '测试3', 75)")
	return err
}

func deleteData(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM sms WHERE number IN (9997, 9998, 9999)")
	return err
}

func TestQueryRow(t *testing.T) {
	db := initDB()
	defer db.Close()
	if err := setup(db); err != nil {
		t.Fatalf("设置测试数据失败: %s", err)
	}
	defer deleteData(db)
	student, err := queryRow(9997)
	if err != nil {
		t.Fatalf("queryRow返回了一个错误: %s", err)
	}
	expectedStudent := Student{Number: 9997, Name: "测试1", Score: 88}
	if student != expectedStudent {
		t.Errorf("预期得到的学生: %v名, 实际得到的学生: %v名", expectedStudent, student)
	}
}

func TestQueryMultiRow(t *testing.T) {
	db := initDB()
	defer db.Close()
	if err := setup(db); err != nil {
		t.Fatalf("设置测试数据失败: %s", err)
	}
	defer deleteData(db)
	students, err := queryMultiRow()
	if err != nil {
		t.Fatalf("queryMultiRow 返回了一个错误: %s", err)
	}
	expectedStudents := map[int]Student{
		9997: {Number: 9997, Name: "测试1", Score: 88},
		9998: {Number: 9998, Name: "测试2", Score: 92},
		9999: {Number: 9999, Name: "测试3", Score: 75},
	}
	found := 0
	for _, student := range students {
		if expStu, ok := expectedStudents[student.Number]; ok {
			if expStu.Name != student.Name || expStu.Score != student.Score {
				t.Errorf("预期得到的学生信息: %v, 实际得到的学生信息: %v", expStu, student)
			}
			found++
		}
	}
	if found != len(expectedStudents) {
		t.Errorf("预期查询到的特定学生数量: %d, 实际查询到的特定学生数量: %d", len(expectedStudents), found)
	}
}

func TestDeleteRow(t *testing.T) {
	db := initDB()
	defer db.Close()
	err := setup(db)
	if err != nil {
		t.Fatalf("设置测试数据失败: %s", err)
	}
	defer func() {
		err := deleteData(db)
		if err != nil {
			t.Errorf("删除测试数据失败: %v", err)
		}
	}()
	number := 9997
	err = deleteRow(number)
	if err != nil {
		t.Errorf("错误为： %v", err)
	}
	numberNotExists := 12345
	err = deleteRow(numberNotExists)
	if err == nil || err.Error() != fmt.Sprintf("没有找到学号为 %d 的学生", numberNotExists) {
		t.Errorf("错误： '没有找到学号为 %d 的学生', 错误为： %v", numberNotExists, err)
	}
}

func TestUpdateRow(t *testing.T) {
	db := initDB()
	defer db.Close()
	err := setup(db)
	if err != nil {
		t.Fatalf("设置测试数据失败: %s", err)
	}
	defer deleteData(db)
	originalStudent, err := queryRow(9997)
	if err != nil {
		t.Fatalf("查询原始学生信息失败: %s", err)
	}
	newScore := 95
	err = updateRow(88, newScore)
	if err != nil {
		t.Fatalf("更新学生信息失败: %s", err)
	}
	updatedStudent, err := queryRow(9997)
	if err != nil {
		t.Fatalf("查询更新后的学生信息失败: %s", err)
	}
	if updatedStudent.Score != newScore {
		t.Errorf("预期得到的分数: %d, 实际得到的分数: %d", newScore, updatedStudent.Score)
	}
	if updatedStudent.Number != originalStudent.Number || updatedStudent.Name != originalStudent.Name {
		t.Errorf("学生的其他信息发生了不期望的变化")
	}
}

func TestInsertRow(t *testing.T) {
	db := initDB()
	defer db.Close()
	err := insertRow(10000, "测试4", 90)
	if err != nil {
		t.Fatalf("测试数据添加错误，错误为：%s\n", err)
	}
	defer func() {
		// 删除测试数据
		err := deleteRow(10000)
		if err != nil {
			t.Errorf("测试数据删除失败，错误为：%s\n", err)
		}
	}()
	stu, err := queryRow(10000)
	if err != nil {
		t.Errorf("查询添加后的测试数据错误，错误为：%s\n", err)
	}
	if stu.Number != 10000 || stu.Name != "测试4" || stu.Score != 90 {
		t.Errorf("测试数据添加失败，请检查失败原因")
	}
}
