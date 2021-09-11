package main

import "fmt"

func main() {

	var num1 int = 200

	var num2 int = 300

	//num1, num2 = max(num1, num2)

	//fmt.Println("", num1, num2)

	fmt.Println("a,b的值是", num1, num2)
	swap(num1, num2)
	fmt.Println("a,b的值是", num1, num2)

	fmt.Println("a,b的值是", num1, num2)
	swap2(&num1, &num2)
	fmt.Println("a,b的值是", num1, num2)

	var a *int

	a = &num1

	fmt.Printf("num1的%v\n", &num1)
	fmt.Printf("a的%v", a)

}
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
引入传递
*/
func swap2(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp

}
