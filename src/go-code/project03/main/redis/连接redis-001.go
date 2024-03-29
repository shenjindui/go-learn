package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//Dial

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")

	defer c.Close()
}
