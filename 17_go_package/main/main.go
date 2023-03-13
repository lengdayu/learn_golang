package main

import (
	"fmt"
	"goToDo/17_go_package/calc"
)

func main() {
	calc.Name = "是是是"
	fmt.Printf("结果是：%v\n", calc.Add(1, 2))
}
