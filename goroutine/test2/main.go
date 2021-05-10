package main

import (
	"fmt"
	"time"
)

func main() {
	task := make(chan int, 10)

	go consumer(task)

	//生产者
	for i := 0; i < 10; i++ {
		task <- i
	}
	time.Sleep(time.Second * 2)
}

//消费者
func consumer(task <-chan int) {
	for i := 0; i < 10; i++ {
		go func(id int) {
			t := <-task
			fmt.Println(id, t)
			time.Sleep(time.Second)
		}(i)
	}
}
