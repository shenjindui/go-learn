package main

import "fmt"

/**
函数学习
*/
func main() {

	num1 := 200
	num2 := 300
	num1, num2 = max(num1, num2)

	//求两个数的最大值
	fmt.Println("", num1, num2)

	//值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	fmt.Println("a,b的值是", num1, num2)
	swap(num1, num2)
	fmt.Println("a,b的值是", num1, num2)
	//引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
	fmt.Println("a,b的值是", num1, num2)
	swap2(&num1, &num2)
	fmt.Println("a,b的值是", num1, num2)

	var a *int
	a = &num1
	fmt.Printf("num1=%v\n", &num1)
	fmt.Printf("a=%v\n", a)

	//判断是否存在
	fmt.Println(multiRet("one"))

	//不定参数求和
	intArray := []int{1, 3, 4, 5, 6}
	sum := getSum(intArray...)
	sum1 := getSum(1, 2, 3)
	fmt.Printf("sum=%v\n", sum)
	fmt.Printf("sum=%v\n", sum1)

	//闭包函数
	nextNumFunc := nextNum()
	//fmt.Printf("type=%T,value=%v",nextNumFunc,nextNumFunc)
	for i := 0; i < 10; i++ {
		fmt.Println(nextNumFunc())
	}

	// 递归函数
	fmt.Println("递归之后的值", fact(4))

}

/**
计算两数的大者  golang函数支持返回多个
*/
func max(num1 int, num2 int) (int, int) {
	if num1 > num2 {
		return num1, num2
	}
	return num2, num1
}

/**
值传递
*/
func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

/**
引用传递
*/
func swap2(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp

}

/**
查找是否存在否个key 返回多个值
*/
func multiRet(key string) (int, bool) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	var err bool
	var val int
	val, err = m[key]
	return val, err
}

/**
  不定参数列表求和 ...指的是不知道参数列表的个数
*/
func getSum(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum = sum + val
	}
	return sum
}

/*
  闭包函数
*/
func nextNum() func() int {
	i, j := 0, 1
	return func() int {
		var tmp = i + j
		i, j = j, tmp
		return tmp
	}
}

/**
  递归函数
*/
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
