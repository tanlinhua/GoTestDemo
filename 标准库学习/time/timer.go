package main

import (
	"fmt"
	"time"
)

func main() {
	go testSleep()
	go testTick()

	time.Sleep(60 * time.Second)
}

func testSleep() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("Sleep", time.Now())
	}
}

func testTick() {
	t1 := time.Tick(2 * time.Second)
	for {
		select {
		case <-t1:
			fmt.Println("Tick", time.Now())
		}
	}
}
