package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var count = 10000

	for task := 0; task < count; task++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
}

func test() {
	defer wg.Done()

	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("http.Get Error", err.Error())
		return
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll Error", err.Error())
		return
	}
	lenght := len(string(body))
	fmt.Println("body.lenght", lenght)
	fmt.Println("Status", resp.Status)
	fmt.Println(runtime.NumGoroutine())
}
