package main

import "fmt"

func main() {

	//Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
	//
	//Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

	//声明变量，默认 map 是 nil  如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

	var maps = make(map[string]string)

	maps["China"] = "beijing"

	for key := range maps { //  range 默认返回的是key 也可以key value一起返回
		fmt.Println(key, maps[key])
	}

	key, ok := maps["China"] //判断是否存在某一个key

	fmt.Println(key, ok)

	if ok {
		fmt.Println(key)
	} else {
		fmt.Println("不存在")
	}

	//delete 删除某一个元素
	fmt.Println(len(maps))
	delete(maps, "China")
	fmt.Println(len(maps))
}
