package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

type studentMgr struct {
	allStudents []*student
}

func newStudent(id int, name, class string) *student {
	return &student{
		id,
		name,
		class,
	}
}

func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 方法：
// 1.添加学生
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2.编辑学生
func (s *studentMgr) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		if v.id == newStu.id {
			s.allStudents[i] = newStu
			return
		}
		fmt.Printf("没有找到学号：%d 对应的学生，请确认输入的学号~\n", newStu.id)
	}
}

// 3.展示所有学生
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d  姓名：%s  班级：%s \n", v.id, v.name, v.class)
	}
}
