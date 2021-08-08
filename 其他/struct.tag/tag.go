package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func GetTag() {
	u := User{}
	ut := reflect.TypeOf(u)
	for i := 0; i < ut.NumField(); i++ {
		name := ut.Field(i).Name
		tag := ut.Field(i).Tag.Get("json")
		fmt.Printf("字段:%s,标签:%s\n", name, tag)
	}
}
