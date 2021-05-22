package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// 通过信号量控制goroutine并发执行的数量
func main() {
	s := semaphore.NewWeighted(10)

	ctx := context.Background()
	for i := 0; i < 20; i++ {
		s.Acquire(ctx, 1)

		go func(i int) {
			defer s.Release(1)
			fmt.Println(i)
			time.Sleep(3 * time.Second)
		}(i)
	}

	// 请求3个资源保证前面的任务都已经执行完毕
	err := s.Acquire(ctx, 10)
	if err != nil {
		fmt.Println(err)
	}
}
