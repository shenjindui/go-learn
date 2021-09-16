package main

import (
	"fmt"
)

type Phone interface {
	call(a int)

	getNameById(id string)
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call(a int) {
	fmt.Println("I am Nokia, I can call you!", a)
}
func (nokiaPhone NokiaPhone) getNameById(id string) {
	fmt.Println("I am Nokia, I can call you!", id)
}

type IPhone struct {
}

func (iPhone IPhone) call(a int) {
	fmt.Println("I am iPhone, I can call you!", a)
}

func (iPhone IPhone) getNameById(id string) {
	fmt.Println("I am iPhone, I can call you!", id)
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call(1)

	phone = new(IPhone)
	phone.call(2)

}
