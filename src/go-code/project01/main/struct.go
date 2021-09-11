package main

import "fmt"

type People struct {
	name string
	age  int
}

func main() {

	var p1 People
	p1.age = 28
	p1.name = "shinnied"

	//fmt.Println(p1.name =="")
	test(&p1)

}

func test(p *People) {
	fmt.Printf("%T %v\n", p.name, p.name)

	fmt.Println("--->", (*p).name)
}
