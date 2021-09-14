package main

import "fmt"

/**
  切片 [] 标识切片类型
*/
func main() {
	//切片 slice  切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。  [] 表示是切片类型
	//切片是可索引的，并且可以由 len() 方法获取长度。 切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

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

	//copy函数
	number1 := make([]int, len(slice3), cap(slice3)*2)
	copy(number1, slice3)
	fmt.Printf("number1 = %v\n", len(number1))
	string := []string{"nihao", "shenjindui", "1000"}
	test1(string...)

	//slice自身维护了一个指针属性，指向它底层数组中的某些元素的集合。 (可以理解为就是一个指针)
	//每一个slice结构都由3部分组成：容量cap(capacity)、长度len(length)和指向底层数组某元素的指针，
	//它们各占8字节(1个机器字长，64位机器上一个机器字长为64bit，共8字节大小，32位架构则是32bit，占用4字节)，所以任何一个slice都是24字节(3个机器字长)
	testSlice := make([]int, 3, 5)
	fmt.Println("testSlice = ", testSlice)

	//make()和new()
	//make()比new()函数多一些操作，new()函数只会进行内存分配并做默认的赋0初始化，而make()可以先为底层数组分配好内存，
	//然后从这个底层数组中再额外生成一个slice并初始化。另外，make只能构建slice、map和channel这3种结构的数据对象，
	//因为它们都指向底层数据结构，都需要先为底层数据结构分配好内存并初始化。

	//array 和 slice的区别
	// 创建长度为3的int数组
	array := [3]int{10, 20, 30}
	fmt.Println("array=", array)
	// 创建长度和容量都为3的slice
	slice1 := []int{10, 20, 30}
	fmt.Println("slice=", cap(slice1))

	//对slice进行切片 (左闭右开)
	mySlice := []int{11, 22, 33, 44, 55}
	// 生成新的slice，从第二个元素取，切取的长度为2
	newSlice := mySlice[1:3]
	fmt.Println("newSlice=", newSlice)

	//合并slice 合并slice时，通过append函数,只需将append()的第二个参数后加上  ...
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	s3 := append(s1, s2...)
	fmt.Println(s3)

	//传递slice给函数
	//Go中函数的参数是按值传递的，所以调用函数时会复制一个参数的副本传递给函数。如果传递给函数的是slice，
	//它将复制该slice副本给函数，这个副本实际上就是[3/5]0xc42003df10，所以传递给函数的副本仍然指向源slice的底层数组。
	//换句话说，如果函数内部对slice进行了修改，有可能会直接影响函数外部的底层数组，从而影响其它slice。但并不总是如此，
	//例如函数内部对slice进行扩容，扩容时生成了一个新的底层数组，函数后续的代码只对新的底层数组操作，这样就不会影响原始的底层数组

	//虽然函数参数传递是拷贝一份，然而拷贝的一份地址仍然指向原来的那一个slice.
	s4 := []int{11, 22, 33, 44}
	foo(s4)
	fmt.Println(s4[1]) // 输出：23

	//内存上关于slice

	//由于slice的底层是数组，很可能数组很大，但slice所取的元素数量却很小，这就导致数组占用的绝大多数空间是被浪费的。
	//特别地，垃圾回收器(GC)不会回收正在被引用的对象，当一个函数直接返回指向底层数组的slice时，这个底层数组将不会随函数退出而被回收，
	//而是因为slice的引用而永远保留，除非返回的slice也消失。
	//因此，当函数的返回值是一个指向底层数组的数据结构时(如slice)，应当在函数内部将slice拷贝一份保存到一个使用自己底层数组的新slice中，
	//并返回这个新的slice。这样函数一退出，原来那个体积较大的底层数组就会被回收，保留在内存中的是小的slice。
}

/**
  测试 ...
*/
func test1(args ...string) {
	for _, v := range args {
		fmt.Printf("%v\n", v)
	}
}

func foo(s []int) {
	for index, _ := range s {
		s[index] += 1
	}
}
