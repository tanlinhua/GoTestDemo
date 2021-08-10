package main

import (
	"fmt"
	"strconv"
)

type Sleeper interface {
	Sleep()
}

type Eater interface {
	Eat(foodName string)
}

type LazyAnimal interface {
	Sleeper
	Eater
}

// Dog 对象:狗
type Dog struct {
	Name string
}

func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping\n", d.Name)
}

func (d Dog) Eat(foodName string) {
	fmt.Printf("Dog %s is eating %s\n", d.Name, foodName)
}

// Cat 对象:猫
type Cat struct {
	Name string
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping\n", c.Name)
}

func (c Cat) Eat(foodName string) {
	fmt.Printf("Cat %s is eating %s\n", c.Name, foodName)
}

func AnimalSleep(s Sleeper) {
	s.Sleep()
}

// Test1 测试 Sleeper 接口
func Test1() {
	var s Sleeper
	dog := Dog{Name: "xiaohei"}
	cat := Cat{Name: "kitty"}

	fmt.Println("第1种调用方法")
	s = dog
	AnimalSleep(s)
	s = cat
	AnimalSleep(s)

	fmt.Println("第2种调用方法")
	sleepList := []Sleeper{Dog{Name: "xiaohei2"}, Cat{Name: "kitty2"}}
	for _, s := range sleepList {
		s.Sleep()
	}
}

// Test2 测试 LazyAnimal 接口
func Test2() {
	sleepList := []LazyAnimal{Dog{Name: "xiaohei3"}, Cat{Name: "kitty3"}}
	foodName := "apple"
	for i, s := range sleepList {
		s.Sleep()
		s.Eat(foodName + strconv.Itoa(i))
	}
}

func main() {
	fmt.Println("Test1")
	Test1()
	fmt.Println("Test2")
	Test2()
}
