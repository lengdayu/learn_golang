package main

import (
	"fmt"
	"os"
)

// 学员管理系统
// 1.添加学员信息
// 2.编辑学员信息
// 3.展示所有学员信息
func showMenu() {
	fmt.Print("\n")
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1.添加学员信息")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出学员信息系统")
	fmt.Print("\n")
}

// 获取用户输入的学员信息
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按要求输入学员信息")
	fmt.Print("请输入学员学号")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员姓名")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员班级")
	fmt.Scanf("%s\n", &class)
	stu := newStudent(id, name, class)
	return stu

}

func main() {
	sm := newStudentMgr()
	for {
		//  1.打印系统菜单
		showMenu()
		//	2.等待用户选择要执行的选项
		fmt.Println("请输入你要操作的序号：")
		var input int
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入的是：", input)
		//	3.执行用户选择的选项
		switch input {
		case 1:
			//添加学员
			newStu := getInput()
			sm.addStudent(newStu)
		case 2:
			//编辑学员
			newStu := getInput()
			sm.modifyStudent(newStu)
		case 3:
			//展示所有学员
			sm.showStudent()
		case 4:
			//退出
			os.Exit(0)
		}
	}

}
