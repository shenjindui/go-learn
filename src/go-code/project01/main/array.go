package main

import "fmt"

func main() {

	// ... 可以理解为不定长的参数。 有多个不定参数的情况，可以接受多个不确定数量的参数。

	var numArray [5]int

	numArray[0] = 5

	numArray1 := [...]float64{1.33, 2.03434321, 32, 32, 32, 32, 32, 32, 32}

	//fmt.Println("第一个值是",numArray[2])

	fmt.Println("第一个值是", numArray1[1])

	//二维数据 append 添加

	var numArray3 [][]int

	num1 := []int{1, 2, 3}
	num2 := []int{1, 2, 3}
	numArray3 = append(numArray3, num1)
	numArray3 = append(numArray3, num2)

	fmt.Println(numArray3[0][0])

	//二位数组初始化
	a := [][]int{
		{
			1, 2, 3,
		},
		{
			4, 5, 6,
		},
	}
	fmt.Println("a", a[0][0])
}
