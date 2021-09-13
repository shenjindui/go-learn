package main

import "fmt"

const MAX = 3

/**
  指针的学习 取址符号 &
*/
func main() {

	//指针数组
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

	//当一个指针被定义后没有分配到任何变量时，它的值为 nil。 nil 指针也称为空指针。
	var pointer *int
	if pointer == nil {
		fmt.Println("the point is =", pointer)
		fmt.Printf("the point value = %x\n", pointer)
	}

	//指向指针的指针.如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
	var a2 int
	var ptr1 *int
	var pptr1 **int
	a2 = 3000
	/* 指针 ptr1 地址 */
	ptr1 = &a2
	/* 指向指针 ptr1 地址 */
	pptr1 = &ptr1
	/* 获取 pptr1 的值 */
	fmt.Printf("变量 a2 = %d\n", a2)
	fmt.Printf("指针变量 *ptr1 = %d\n", *ptr1)
	fmt.Printf("指向指针的指针变量 **pptr1 = %d\n", **pptr1)
	fmt.Printf("变量的地址 a2 = %d\n", &a2)
	fmt.Printf("指针变量 *ptr1的地址 = %d\n", ptr1)
	fmt.Printf("指向指针的指针变量 **pptr1的地址 = %d\n", pptr1)

	//指针作为函数参数
	/* 定义局部变量 */
	a4 := 100
	b := 200
	fmt.Printf("交换前 a4 的值 : %d\n", a4)
	fmt.Printf("交换前 b 的值 : %d\n", b)
	/* 调用函数用于交换值
	 * &a4 指向 a4 变量的地址
	 * &b 指向 b 变量的地址
	 */
	swap4(&a4, &b)
	fmt.Printf("交换后 a4 的值 : %d\n", a4)
	fmt.Printf("交换后 b 的值 : %d\n", b)

	//关于传值问题   在Go语言里，所有的参数传递都是值传递(传值)，都是一个副本，一个拷贝

	//函数返回变量的指针时，这个变量会逃逸
	//当觉得栈上的空间不够时，会分配在堆上
	//在切片上存储指针或带指针的值的时候，对应的变量会逃逸
	//chan里面的元素是指针的时候，也会发生逃逸
	//map的value是指针的时候，也会发生逃逸
	//在interface类型上调用方法，也会发生逃逸
	//当给一个slice分配一个动态的空间容量时，也会发生逃逸
	//函数或闭包外声明指针，在函数或闭包内分配，也会发生逃逸
	//函数外初始化变量，函数内使用变量，然后返回函数，也会发生逃逸
	//被已经逃逸的指针引用的指针，也会发生逃逸
	//逃逸分析在编译阶段完成的
	//逃逸分析的好处可以判断变量放在堆上还是栈上，分配在栈上的变量对GC友好

}

/**
  使用指针交换两个数
*/
func swap4(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
