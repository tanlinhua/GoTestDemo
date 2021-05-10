package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"pprof.test/data"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/EDDYCJY"))
			time.Sleep(time.Second * 2)
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}

// https://www.jianshu.com/p/4e4ff6be6af9
