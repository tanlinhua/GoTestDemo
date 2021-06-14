package main

import "fmt"

type Factory interface {
	Method1() //处理方法1
}

//奶牛
type Cow struct {
	Name string
}

//奶牛的处理方法，实现接口
func (cow Cow) Method1() {
	fmt.Printf("%s：挤奶\n", cow.Name)
}

//苹果
type Apple struct {
	Name string
}

//苹果的处理方法，实现接口
func (a Apple) Method1() {
	fmt.Printf("%s：切块\n", a.Name)
}

func main() {

	// var fa Factory
	// cow := Cow{"奶牛"}
	// fa = cow
	// fa.Method1()

	// apple := Apple{"苹果"}
	// fa = apple
	// fa.Method1()

	falg := 2
	var fa Factory

	if falg == 1 {
		fa = Cow{"Cow"}
	} else {
		fa = Apple{"苹果"}
	}
	fa.Method1()
}
