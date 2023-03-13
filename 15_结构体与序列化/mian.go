package main

import (
	"encoding/json"
	"fmt"
)

//结构体字段的可见性与JSON序列化

// 如果一个Go语言包中定义的标识符的首字母是大写的，那么对外就是可见的
type student struct {
	ID   int
	Name string
}

type class struct {
	Title    string    `json:"title"`
	Students []student `json:"students"`
}

func NewStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

func main() {
	//创建一个班级变量c1
	c1 := class{
		Title:    "火箭101",
		Students: make([]student, 0, 20),
	}
	for i := 0; i <= 10; i++ {
		//	造10个学生
		c1.Students = append(c1.Students, NewStudent(i, fmt.Sprintf("stu%02d", i)))
	}
	//fmt.Printf("%v\n", c1)
	//JSON序列化 Go语言中的数据-> JSON的数据格式
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed ,err:", err)
		return
	} else {
		fmt.Print("\n")
		fmt.Printf("%T\n", data)
		fmt.Printf("%s\n", data)

	}
	jsonStr := fmt.Sprintf("%s", data)
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json unmarshall failed err: \n", err)
		return
	} else {
		fmt.Printf("%#v\n", c2)
	}
}
