package main

import "fmt"

/**
  介绍for循环
*/
func main() {

	//1.普通的for循环 ----单个for循环
	//init： 一般为赋值表达式，给控制变量赋初值； condition： 关系表达式或逻辑表达式，循环控制条件； post： 一般为赋值表达式，给控制变量增量或减量。
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("%T %v\n", sum, sum)

	//2.使用range  字符串string、数组array、切片slice,map等进行迭代输出元素。
	//range 返回对应的下标(index/key)和值(value),若是我们只想获取值，不需要index/key 可以使用'_' 来代表忽略
	//空白符"_"省略 函数的返回值的使用场景
	//第一种 ：忽略返回值
	//第二种：用在变量
	//第三种：用在import package  执行import代码之前会先调用test/foo中的初始化函数(init)，这种使用方式仅让导入的包做初始化，而不使用包中其他功能
	strings := []string{"hello", "golang"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	//2.双重for循环也就是在for循环之内在嵌套一层for循环
	//九九乘法表
	for m := 1; m < 10; m++ {
		/*    fmt.Printf("第%d次：\n",m) */
		for n := 1; n <= m; n++ {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
		}
		fmt.Println("")
	}

	//3.while do-while

	//4.循环控制语句 break continue goto 其中记录一下goto
	//Go 语言的 goto 语句可以无条件地转移到过程中指定的行。可用来实现条件转移， 构成循环，跳出循环体等功能。但是我们平时开发时不建议使用
	gotoTag()

}

//for循环配合goto打印九九乘法表
func gotoTag() {
	for m := 1; m < 10; m++ {
		n := 1
	LOOP:
		if n <= m {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
			n++
			goto LOOP
		} else {
			fmt.Println("")
		}
		n++
	}
}
