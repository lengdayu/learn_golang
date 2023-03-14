package main

import "fmt"

/*
两个goroutine
1.生成0~100的数字发送到ch1
2.从ch1中取出数据计算他的平方，结果发送到ch2中
*/
func f1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
func f2(ch1 <-chan int, ch2 chan<- int) {
	//从通道取值方式一
	for {
		temp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- temp * temp
	}
	close(ch2)
}
func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go f1(ch1)
	go f2(ch1, ch2)
	//	从通道取值方式2
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
