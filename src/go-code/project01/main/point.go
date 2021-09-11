package main

import "fmt"

const MAX = 3

func main() {

	a := []int{1, 2, 3}

	points := [MAX]*int{}

	var i int

	for i = 0; i < MAX; i++ {
		points[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Println("--->", *points[i])
	}

	//指向指针的指针

	var a1 int = 10

	var ptr *int

	var ptr2 **int

	ptr = &a1

	ptr2 = &ptr

	fmt.Println("a-->", &a1)

	fmt.Println("a-->", ptr)

	fmt.Println("a-->", *ptr2)

	var a3 = 3
	var b3 = 4
	fmt.Println("9-->", a3, b3)
	a3, b3 = b3, a3
	fmt.Println("9-->", a3, b3)
}
