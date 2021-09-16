package main

import "fmt"

/////defer recover panic
/**
  panic   类似其他语言throw抛出异常一样，通过函数调用链，一层层的把异常往上抛出去，如果没有人拦截异常，就会退出程序。
          一般在没有recover的情况下panic会导致程序崩溃，panic，defer和recover经常同时出现，用于异常处理。
*/
func main() {
	//1.直接panic
	//testPanic()

	//2.recover 类似try/catch
	//PanicRecoverTest()
	//fmt.Println("------------>")

	//3.
	defer testPanic1()
	panic("出错测试")
	fmt.Println("出错之后")
}
func testPanic1() int {
	error := recover()
	fmt.Println("error", error)
	if error != nil {
		fmt.Println("出错了", recover())
		return 1
	}
	return 0
}

func testPanic() {
	fmt.Println("-----")
	panic("出错了")
}

// 通过defer和recover实现拦截panic错误
//* defer要在panic之前先注册
func PanicRecoverTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("拦截到错误:", err)
		}
	}()
	panic("panic 错误")
	fmt.Println("c")
}
