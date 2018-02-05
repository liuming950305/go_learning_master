package main

import "fmt"

type Student struct {
	studentId int // 学号
	studentName string
	age int
}

func main() {
	var student Student
	var student1 Student

	// 设置值
	student.studentId = 1001
	student.studentName = "大明"
	student.age = 23

	student1.studentId = 1002
	student1.studentName = "小明"
	student1.age = 12

	fmt.Printf("Student studentId: %d\n", student.studentId)
	fmt.Printf("Student studentName: %s\n", student.studentName)
	fmt.Printf("Student age: %d\n", student.age)

	fmt.Printf("Student1 studentId: %d\n", student1.studentId)
	fmt.Printf("Student1 studentName: %s\n", student1.studentName)
	fmt.Printf("Student1 age: %d\n", student1.age)
}


