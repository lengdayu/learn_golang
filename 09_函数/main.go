package main

import "fmt"

// 函数
func go1() {
	fmt.Println("hello world")
}

func intSum(x, y int) int {
	return x + y
}

// 可变参数
// 可变参数在函数中是切片类型
func go2(a ...int) {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

// defer 语句
func go3() {
	print("go3\n")
}

func go4() {
	print("go4\n")
}

// 函数进阶 变量作用域
var num = 10

func go5() {
	fmt.Println("全局变量", num)
}

func main() {
	go1()
	sum1 := intSum(1, 3)
	fmt.Println(sum1)
	go2(1, 2, 3, 4, 5)
	defer go3()
	defer go4()
	go5()
}
