package module

import "database/sql"

type Student struct {
	Id    sql.NullInt64
	Name  sql.NullString
	Score sql.NullInt64
}

var stu = new(Student)

func NewStudent(id int, name string, number int) *Student {
	return &Student{
		Id:    sql.NullInt64{Int64: int64(id), Valid: true},
		Name:  sql.NullString{String: name, Valid: true},
		Score: sql.NullInt64{Int64: int64(number), Valid: true},
	}
}
