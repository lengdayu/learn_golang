package main

import "fmt"

type Animal struct {
	name string
}

type Dog struct {
	feet int8
	*Animal
}

func (d *Dog) Wang() {
	fmt.Printf("%v会汪汪汪\n", d.name)
}
func (d *Dog) Move() {
	fmt.Printf("%v有%v个脚\n", d.name, d.feet)
}
func main() {
	p := &Dog{
		feet:   4,
		Animal: &Animal{name: "乐乐"},
	}
	p.Wang()
	p.Move()
}
