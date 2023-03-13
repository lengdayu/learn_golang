package main

import (
	"fmt"
	"reflect"
)

// 结构体反射
type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

func main() {
	stu1 := student{
		Name:  "乐乐",
		Score: 90,
	}

	//	通过反射去获取结构体中的所有字段
	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%s kind:%v\n", t.Name(), t.Kind())
	//遍历结构体变量所有的字段
	for i := 0; i < t.NumField(); i++ {
		//	根据结构体字段的索引去去字段
		filedObj := t.Field(i)
		fmt.Printf("name:%v type:%v tag:%v \n", filedObj.Name, filedObj.Type, filedObj.Tag)
		fmt.Println(filedObj.Tag.Get("json"), filedObj.Tag.Get("ini"))
	}
	//	根据名字去取结构体中的字段
	filedObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v type:%v tag:%v \n", filedObj2.Name, filedObj2.Type, filedObj2.Tag)
	}
}
