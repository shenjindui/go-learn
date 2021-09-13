package main

import (
	"fmt"
	"sort"
)

/**
  map 学习记录
*/
func main() {
	//Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
	//Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
	//声明变量，默认 map 是 nil  如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	//此时进行取值，返回的是对应类型的零值（不存在也是返回零值）默认值

	//通过关键字make创建map  key的类型 值value的类型
	var maps = make(map[string]string)

	//新增
	maps["China"] = "beijing"

	// 循环遍历 使用 range 默认返回的是key 也可以key value一起返回
	for key := range maps {
		fmt.Println(key, maps[key])
	}

	//判断是否存在某一个key ,返回value的是否存在标识ok
	value, ok := maps["China"]
	fmt.Println(value, ok)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}

	//长度 len()
	fmt.Println(len(maps))
	//delete 删除map某一个元素

	delete(maps, "China")
	fmt.Println(len(maps))

	//对map进行排序
	var testMap = make(map[string]string)
	testMap["f"] = "shinnied"
	testMap["d"] = "checking"
	for _, value := range testMap {
		fmt.Println(value)
	}
	sortMap(testMap)

	//函数传参 map本身是引用类型，作为形参或返回参数的时候，传递的是值的拷贝，而值是地址，扩容时也不会改变这个地址。(传递的是地址值的拷贝)
	var m map[int64]int64
	m = make(map[int64]int64, 1)
	fmt.Printf("m 原始的地址是：%p\n", m)
	changeM(m)
	fmt.Printf("m 改变后地址是：%p\n", m)
	fmt.Println("m 长度是", len(m))
	fmt.Println("m 参数是", m)

	//map的数据结构理解 hmap sync gc
}

/**
测试对map根据key进行排序
*/
func sortMap(m map[string]string) {
	var keys []string
	// 把key单独抽取出来，放在数组中
	for k, _ := range m {
		keys = append(keys, k) //append对数组添加元素
	}
	// 进行数组的排序
	sort.Strings(keys)
	// 遍历数组就是有序的了
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}

// 改变map的函数
func changeM(m map[int64]int64) {
	fmt.Printf("m 函数开始时地址是：%p\n", m)
	var max = 5
	for i := 0; i < max; i++ {
		m[int64(i)] = 2
	}
	fmt.Printf("m 在函数返回前地址是：%p\n", m)
}
