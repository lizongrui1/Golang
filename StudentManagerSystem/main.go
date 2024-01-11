package main

import (
	"fmt"
	"http-sql/module"
)

func main() {
	err := module.InitDB()
	if err != nil {
		fmt.Printf("Failed to initialize DB: %v\n", err)
		return
	}

	student := module.NewStudent(6, "xiaoli", 10086)
	fmt.Printf("Student ID: %d, Name: %s, Number: %d\n", student.Id.Int64, student.Name.String, student.Number.Int64)
}
