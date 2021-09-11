package main

import "fmt"

func main() {
	var num int = 3

	var string string = "你好"

	switch num {
	case 6:
		fmt.Printf("%T 值是%v", num, num)
	case 3:
		fmt.Printf("%T 值是%v\n", num, num)
	}

	switch string {
	case "shinnied":
		fmt.Printf("%T 值是%v\n", string, string)
		fallthrough
	case "你好1", "你好2":
		fmt.Printf("%T 值是%v", string, string)
	default:
		fmt.Printf("%T 值是%v", string, string)
	}

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("%T", i)
	case int32:
		fmt.Printf("%T", i)

	}
}
