package main

import "fmt"

func main() {
	var score = 35
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 60 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	for i := 0; i <= 10; i++ {
		// fmt.Println(i)
	}
	var i1 uint = 1
	i2 := 1
	fmt.Println(i1)
	fmt.Println(i2)
	var i int = 3
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}
