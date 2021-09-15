package main

import (
	"fmt"
	"time"
)

/**
  并发 通道 channel
*/
func main() {

	//1.Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。 goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的
	//Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。
	//****同一个程序中的所有 goroutine 共享同一个地址空间。
	go say("world")
	say("hello")
	//输出的 hello 和 world 是没有固定先后顺序。因为它们是两个 goroutine 在执行。

	//2.channel 通道 说白就是为了再两个goroutine进行通讯和同步      a<-b   b<-a
	//通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
	//ch <- v    // 把 v 发送到通道 ch
	//v := <-ch  // 从 ch 接收数据 并把值赋给 v

	//3.创建通道 make  (默认情况下，通道是不带 缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。)
	ch := make(chan int) // make(chan int,0)
	//当capacity=0时，channel是无缓冲阻塞读写的，当capacity>0时，channel是有缓冲非阻塞的，直到写满capacity元素才阻塞写入
	fmt.Printf("%T\n", ch)
	s := []int{7, 2, 8, -9, 4, 0}
	fmt.Println("数组", s[:len(s)/2])
	c := make(chan int)
	go sum(s[:len(s)/2], c) //17
	go sum(s[len(s)/2:], c) //-5
	//x := <-c
	//fmt.Println(x)
	x, y := <-c, <-c // 从通道 c 中接收 (先进先出)//本质上是一个队列 先进先出 channel本身就是线程安全
	fmt.Println(x, y, x+y)

	ch2 := make(chan int)
	//ch2 <-2 //这里不能写的原因是 main函数被阻塞了，此时不能在继续往下执行了，
	//x2 := <-ch2
	//fmt.Println(x2)
	//可以开启两个goroutine
	go testW(ch2)
	go testR(ch2)
	time.Sleep(time.Millisecond)

	//4.通道缓冲区 通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小
	//带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于(异步)状态，就是说发送端发送的数据可以放在(缓冲区)里面，
	//可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。

	// 缓冲区大小为2
	che := make(chan int, 2)
	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	che <- 1
	che <- 2
	// 获取这两个数据
	fmt.Println(<-che)
	fmt.Println(<-che)

	//5.关闭通道close()
	ce := make(chan int, 10)
	go fibonacci(cap(ce), ce)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range ce {
		fmt.Println("value=", i)
	}

	testChan := make(chan int, 1)
	testChan <- 0
	i := <-testChan
	fmt.Println(i, "--->")
	testChan <- 1
	i = <-testChan
	fmt.Println(i, "--->")
	//close(testChan)
	//strings := <-testChan
	//fmt.Println(strings, "--->")
	//testChan<-"shinnied"
}

func testW(ch chan int) {
	ch <- 4
	ch <- 6
}
func testR(ch chan int) {
	for {
		v, ok := <-ch
		fmt.Println(v, ok)
	}
}

/**
  goroutine测试
*/
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/**
求和
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

/**
  关闭通道
*/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
