package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("%T %v\n", sum, sum)

	strings := []string{"Miho", "你好"}

	for i, s := range strings {
		fmt.Println("", i, s)
	}
}
