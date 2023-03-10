package main

import "fmt"

func main() {
	fmt.Println("class")
	go1()
	result := newPerson("zhuzhu", "上海", 30)

	fmt.Printf("%#v\n", result)
}

type MyInt int
type NewType = int

func go1() {
	//	自定义类型和类型别名示例
	var i MyInt
	fmt.Printf("type:%T, value:%v \n", i, i)

	//	类型别名
	var x NewType
	fmt.Printf("type:%T, value:%v \n", x, x)

	go2()

}

type person struct {
	name, city string
	age        int8
}

func go2() {
	p := person{
		name: "xiaoMing",
		city: "江苏",
		age:  18,
	}

	fmt.Printf("xiaoming: %T,%#v\n", p, p)

	p2 := new(person)
	p2.name = "小可爱"
	p2.city = "上海"

	fmt.Printf("p2:%#v\n", p2)

	p3 := &person{}
	fmt.Printf("%#v\n", p3)
	fmt.Printf("%T\n", p3)
	p4 := person{
		"小军",
		"南京",
		27,
	}
	fmt.Printf("%#v\n", p4)
	fmt.Printf("%T\n", p4)

}

// 结构体构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name,
		city,
		age,
	}
}
