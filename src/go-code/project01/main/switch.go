package main

import "fmt"

/**
  switch的学习记录
*/
func main() {
	var num int = 3

	var string1 string = "你好"
	//switch 的用法跟其他语言差不多，但是分支后面不需要加 break; case 可以是多表达式(支持多条件匹配)
	switch num {
	case 6:
		fmt.Printf("%T 值是%v", num, num)
	case 3:
		fmt.Printf("%T 值是%v\n", num, num)
	}

	switch string1 {
	case "shinnied":
		fmt.Printf("%T 值是%v\n", string1, string1)
		fallthrough
	case "你好1", "你好2":
		fmt.Printf("%T 值是%v\n", string1, string1)
	default:
		fmt.Printf("%T 值是%v\n", string1, string1)
	}

	//switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T\n", i)
	case int:
		fmt.Printf("x 是 int 型\n")
	case float64:
		fmt.Printf("x 是 float64 型\n")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型\n")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型\n")
	default:
		fmt.Printf("未知型\n")
	}

	//fallthrough  使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	//如果想要执行多个 case，需要使用 fallthrough 关键字，也可用 break 终止。
	testNum := 2
	switch testNum {
	case 2:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case 3:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case 4:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case 5:
		fmt.Println("4、case 条件语句为 true")
	case 6:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
