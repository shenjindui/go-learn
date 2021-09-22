package main

import (
	"fmt"
)

type User2 struct {
	name string
	age  int
}

type Phone interface {
	call(a int)

	getNameById(id string) User2
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call(a int) {
	fmt.Println("I am Nokia, I can call you!", a)
}
func (nokiaPhone NokiaPhone) getNameById(id string) User2 {
	return User2{
		name: "shinnied",
		age:  25,
	}
}

type IPhone struct {
}

func (iPhone IPhone) call(a int) {
	fmt.Println("I am iPhone, I can call you!", a)
}

func (iPhone IPhone) getNameById(id string) User2 {
	fmt.Println("I am iPhone, I can call you!", id)
	return User2{}
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call(1)
	result := phone.getNameById("323232")
	fmt.Println(result.name)

	phone = new(IPhone)
	phone.call(2)

	var x interface{}
	x = 100
	v, ok := x.(int)
	fmt.Println(x, ok)
	if ok {
		fmt.Println(v)
	}

}
