package main

import "fmt"

//使用值接收者和使用指针接收者的区别

type mover interface {
	move()
}
type person struct {
	name string
	age  int8
}

func (p person) move() {
	fmt.Printf("%s在跑..\n", p.name)
}

//使用值接收者实现接口

func main() {
	var m mover
	var m2 mover
	p1 := person{
		name: "猪猪",
		age:  18,
	}
	m = p1
	m.move()
	fmt.Println(m)

	p2 := &person{
		name: "乐乐",
		age:  22,
	}
	m2 = p2
	m2.move()
	fmt.Println(m2)
}
