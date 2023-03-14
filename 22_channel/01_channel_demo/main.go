package main

import "fmt"

func main() {
	var ch1 chan int
	ch1 = make(chan int, 1) //缓冲区最大存储一个数值
	ch1 <- 10               //发送值
	x := <-ch1              //接收值
	fmt.Println(x)
	close(ch1)
}
