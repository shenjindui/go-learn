package main

import "fmt"

func main() {
	//切片 slice  切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。  [] 表示是切片类型
	//切片是可索引的，并且可以由 len() 方法获取长度。
	//切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
	//一个切片在未初始化之前默认为 nil，长度为 0
	var slice []int = make([]int, 5, 5)

	slice[0] = '2'
	//fmt.Println(slice[0])

	var slice2 []int

	if slice2 == nil {
		//fmt.Printf("kongde ")
	}

	//fmt.Printf("%T\n %v",slice2,slice2)

	//增加切片的容量 使用append copy 函数

	var slice3 []int

	fmt.Printf("slice3%v\n", slice3)

	slice3 = append(slice3, 0)

	fmt.Printf("slice3%v\n", slice3)

	slice3 = append(slice3, 1, 2, 3, 4)

	fmt.Printf("slice3%v\n", slice3)

	fmt.Printf("slice3 cap = %v\n", cap(slice3))

	number1 := make([]int, len(slice3), cap(slice3)*2)

	copy(number1, slice3)

	fmt.Printf("number1 = %v\n", len(number1))

	string := []string{"nihao", "shenjindui", "1000"}
	test1(string...)
}

/**
  ...
*/
func test1(args ...string) {
	for _, v := range args {
		fmt.Printf("%v\n", v)
	}
}
