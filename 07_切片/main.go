package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3}
	go1(a)
	go2()
}

func go1(x [3]int) int {
	sum := 0
	for _, value := range x {
		fmt.Println(value)
		sum += value
	}

	return 1
}
func go2() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]
	fmt.Println(s)
	fmt.Println(len(s)) //长度
	fmt.Println(cap(s)) //容量
}
