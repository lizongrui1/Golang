package module

type Student struct {
	Id    int
	Name  string
	Score int
}

var stu = new(Student)

func NewStudent(id int, name string, number int) *Student {
	return &Student{
		Id:    id,
		Name:  name,
		Score: number,
	}
}
