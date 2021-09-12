package main

import "fmt"

/**
  数组 array
*/
func main() {

	//1.声明数组  var variable_name [SIZE] variable_type

	var numArray [5]int
	numArray[0] = 5
	// ... 3个点代表不定长的数组，它会自己去推算数组的长度。
	numArray1 := [...]float64{1.33, 2.03434321, 32, 32, 32, 32, 32, 32, 32}
	//fmt.Println("第一个值是",numArray[2])
	fmt.Println("第一个值是", numArray1[1])

	//2.二维数据 append 添加

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

	//3.向函数传递数组
	array := []int{1, 2, 3}
	testArray(array)

	//4.指向数组的指针   和   指针数组
	x, y := 1, 2
	c := [...]*int{&x, &y} //保存的元素是指向int型的指针
	d := [...]int{99: 1}
	var p *[100]int = &d //指向数组的指针，取的是d的地址
	//加*代表这是一个指针
	fmt.Println(p) //打印的结果和d的结果是一样的，只不过前面多了一个取地址的符号 &，这就是指向数组的指针
	fmt.Println(c) //指针数组，存的是x和y的内存地址

	//5.值的比较  数组在Go中为值类型  数组之间可以使用 == 或 != 进行比较，但不可以使用<或>

	//6.可以使用new来创建数组，此方法返回一个指向数组的指针
	newArray := new([10]int) //使用new关键字来创建这个数组的话，它返回的就是这个数组的指针
	newArray[0] = 2
	fmt.Println((*newArray)[0])

	//7.冒泡排序
	bubbleSort()
}

/**
测试往函数参数传递数组
*/
func testArray(array []int) {
	fmt.Printf("类型是=%T,数组的长度是=%v\n", array, len(array))
}

/**
  go的冒泡排序
*/
func bubbleSort() {
	a := [...]int{5, 2, 3, 6, 7, 8, 9}
	fmt.Println(a)

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
