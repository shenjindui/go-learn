package main

import "fmt"

/**
  range 的用法
*/
func main() {
	// range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素
	//在数组和切片中它返回元素的索引和索引对应的值(index - value)，在集合中返回 key-value 对。
	//空白符"_"省略 函数的返回值
	// 第一种 ：忽略返回值
	// 第二种：用在变量
	// 第三种：用在import package  执行import代码之前会先调用test/foo中的初始化函数(init)，这种使用方式仅让导入的包做初始化，而不使用包中其他功能

	slice := []int{1, 2, 3}
	sum := 0
	for _, num := range slice {
		//fmt.Print("v-", i)
		sum = sum + num
	}
	//fmt.Println("num total=", sum)

	for _, value := range "A" {
		fmt.Println(value)
	}

	//map 的初始化  map[类型] 类型

	maps := map[string]string{"a": "hello"}

	for key, value := range maps {
		fmt.Println(key, value)
	}

	array := [3]int{1, 2, 3}

	for _, value := range array {
		fmt.Printf("%T %v", value, value)
	}
}
