package main

/**
  结构体学习 struct
*/
import "fmt"

//声明方式
type People struct {
	name string
	age  int
}
type Info struct {
	sex     int
	name    string
	age     int
	address string
}

type User struct {
	like string
	Info
}

type Admin struct {
	unlike string
	Info
}

//内嵌结构体
type Person1 struct {
	Name  string
	Age   int
	Email string
}
type Student struct {
	Person1
	StudentID int
}

func main() {

	//1.声明并赋值
	var p1 People
	p1.age = 28
	p1.name = "shinnied"
	//fmt.Println(p1.name =="")
	test(&p1)

	//2.匿名结构体
	p2 := struct {
		name string
		age  int
	}{name: "shinnied", age: 25}
	fmt.Println("p2=", p2.age)

	//go语言中虽然没有继承,但是可以 结构内嵌 ,达到类似继承的效果
	user := User{}
	user.sex = 0
	user.address = "广州市"
	user.like = "游戏"
	fmt.Println(user)
	admin := Admin{Info: Info{sex: 1}} //还可以这样声明一些属性值,因为Info是结构体,匿名,所以需要这样声明
	admin.address = "广州市"
	admin.unlike = "游戏"
	fmt.Println(admin)

	//使用new() 函数创建,返回的是一个指针
	//new 是一个用来分配内存的内置函数，但是与 C++ 不一样的是，它并不初始化内存，只是将其置零。也就是说，
	//new(T) 会为类型 T 分配被置零的内存，并且返回它的地址，一个类型为 *T 的值。在 Golang 的术语中，
	//其返回一个指向新分配的类型为 T 的指针，这个指针指向的内容的值为零(zero value)
	nick := new(People)
	(*nick).name = "sjd"
	fmt.Println("nick", nick)

	var nick2 *People
	nick2 = new(People)
	nick2.name = "sjd"
	fmt.Println("nick2", nick2)
	//表达式 new(Type) 和 &Type{} 是等价的。

	//内嵌结构体
	st := Student{
		Person1:   Person1{"jack", 22, "jack@xxx.com"},
		StudentID: 1000,
	}
	fmt.Println("st", st)
}

/**
函数结构体传值   函数中结构体作为参数,如果不是用结构指针,函数内参数属性的改变不影响原来对象的属性的改变
*/
func test(p *People) {
	fmt.Printf("%T %v\n", p.name, p.name)
	fmt.Println("--->", (*p).name)
}
