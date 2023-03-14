package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello goroutine", i)
	wg.Done() //通知wg计数器-1
}

func a() {
	for i := 0; i < 1000; i++ {
		fmt.Println("A:", i)
	}
	wg.Done() //通知wg计数器-1
}
func b() {
	for i := 0; i < 1000; i++ {
		fmt.Println("B:", i)
	}
	wg.Done() //通知wg计数器-1
}
func main() {
	//wg.Add(10000) //计数牌+1
	//for i := 0; i < 10000; i++ {
	//	go hello(i)
	//}
	//fmt.Println("main goroutine done!")
	//wg.Wait() //等待线程执行完毕（计数器为0）

	runtime.GOMAXPROCS(6)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
