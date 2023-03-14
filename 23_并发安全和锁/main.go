package main

import (
	"fmt"
	"sync"
	"time"
)

// 多个goroutine并发操作全局变量

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex   //互斥锁
	rwlock sync.RWMutex //读写互斥锁
)

func add() {
	for i := 0; i < 5000000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

func read() {
	rwlock.Lock()
	time.Sleep(time.Millisecond)
	rwlock.Unlock()
	wg.Done()
}

func write() {
	rwlock.Lock()
	time.Sleep(time.Millisecond * 10)
	rwlock.Unlock()
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)

	//	读写互斥锁demo
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
