package main

import "fmt"

//方法的定义实例

type Person struct {
	name string
	age  int8
}

// NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream 定义方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好GO语言！ \n", p.name)
}

// SetAge 修改年龄的方法
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}
func main() {
	p1 := NewPerson("小菊", 127)
	p1.Dream()
	p1.SetAge(22)
	fmt.Println(p1)
}
