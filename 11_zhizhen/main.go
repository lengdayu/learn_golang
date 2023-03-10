package main

import "fmt"

// 指针
func main() {
	a := 10
	b := &a
	fmt.Printf("a: %v %p\n", a, &a)
	fmt.Printf("b: %v %p\n", b, b)
	//*根据内存地址取值
	c := *b
	fmt.Printf("c:%v %p\n", c, &c)

	d := 1
	modify1(d)
	fmt.Printf("d:%v\n", d)

	e := &d
	modify2(e)
	fmt.Printf("d:%v\n", d)

	newInt()
}

func modify1(x int) {
	x = 100
}
func modify2(y *int) {
	*y = 100
}

func newInt() {
	var a *int
	a = new(int)
	fmt.Println(*a)

}
