package main

import (
	"fmt"
	"sync"
	"time"
)

type Resolve func(resp interface{})
type Reject func(err error)

type PromiseFunc func(resolve Resolve, reject Reject)

type Promise struct {
	f       PromiseFunc
	resolve Resolve
	reject  Reject
	wg      sync.WaitGroup
}

func NewPromise(f PromiseFunc) *Promise {
	return &Promise{f: f}
}

func (p *Promise) Then(resolve Resolve) *Promise {
	p.resolve = resolve
	return p
}

func (p *Promise) Catch(reject Reject) *Promise {
	p.reject = reject
	return p
}

func (p *Promise) Done() {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		p.f(p.resolve, p.reject)
	}()
	p.wg.Wait()
}

// go模拟js promise
func main() {
	NewPromise(func(resolve Resolve, reject Reject) {
		// 模拟业务处理
		if time.Now().Unix()%2 == 0 {
			resolve("OK")
		} else {
			reject(fmt.Errorf("my error"))
		}
	}).Then(func(resp interface{}) {
		fmt.Println(resp)
	}).Catch(func(err error) {
		fmt.Println(err)
	}).Done()
}
