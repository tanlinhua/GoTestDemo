package main

import (
	"errors"
	"fmt"

	"try.catch/exception"
)

// 开启事务
func beginTransaction() {
	fmt.Println("开启事务")
}

// 回滚事务
func rollback() {
	fmt.Println("回滚事务")
}

// 提交事务
func commit() {
	fmt.Println("提交事务")
}

// 执行one操作
func one() (err error) {
	fmt.Println("one ok")
	return nil
}

// 执行two操作
func two() (err error) {
	fmt.Println("two ok")
	return nil
}

// 执行three操作
func three(i int) (err error) {
	if i != 1 {
		panic("three.panic")
	}
	fmt.Println("three ok")
	return nil
}

// 执行four操作
func four() (err error) {
	fmt.Println("four ok")
	return nil
}

// 执行five操作
func five(i int) (err error) {
	if i != 1 {
		err = errors.New("five panic")
		return err
	}
	fmt.Println("five ok")
	return nil
}

func main() {
	var err error
	exception.Block{
		Try: func() {
			beginTransaction()
			if err = one(); err != nil {
				panic(err)
			}
			if err = two(); err != nil {
				panic(err)
			}
			if err = three(1); err != nil {
				panic(err)
			}
			if err = four(); err != nil {
				panic(err)
			}
			if err = five(1); err != nil {
				panic(err)
			}
			err = nil
			commit()
		},
		Catch: func(e interface{}) {
			rollback()
			fmt.Printf("%v panic\n", e)
			err = fmt.Errorf("%v", e)
		},
		Finally: func() {
			fmt.Println("finally")
		},
	}.Do()
	fmt.Println("main.err:", err)
}
