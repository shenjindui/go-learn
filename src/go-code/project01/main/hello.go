package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hello")

	var a int = 4

	var str string
	str = fmt.Sprintf("%d", a)

	fmt.Printf("str type %T  str=%q\n", str, str)

	var a1 = 4
	var str1 string
	str1 = strconv.FormatInt(int64(a1), 10)
	fmt.Printf("%T %q\n", str1, str1)

	var str3 string = "3232"

	var resultInt int64

	var resultInt1 int32

	resultInt, _ = strconv.ParseInt(str3, 10, 64)

	resultInt1 = int32(resultInt)

	fmt.Printf("type is %T value = %v\n", resultInt, resultInt)

	fmt.Printf("type is %T value = %v\n", resultInt1, resultInt1)

	fmt.Printf("%v\n", &resultInt)

	var prt *int64 = &resultInt

	fmt.Printf("%v", prt)

}
