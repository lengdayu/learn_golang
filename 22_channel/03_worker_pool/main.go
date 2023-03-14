package main

import (
	"fmt"
	"time"
)

// worker pool
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("worker:%d stop job:%d\n", id, job)
	}
}
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启多个goroutine
	for j := 0; j < 10; j++ {
		go worker(j, jobs, results)
	}
	// 发送多个任务
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < 100; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}
