package main

import (
	"fmt"
	"sort"
)

func main() {
	var countmap map[string]string
	countmap = make(map[string]string)
	countmap["a"] = "a"
	countmap["b"] = "b"
	for item := range countmap {
		fmt.Println(item)
	}

	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[string]int{"1": 1, "2": 2, "3": 3}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	if map1 == nil {
		map1 = make(map[int]string, 0)
		fmt.Println(map1 == nil)
	}

	map1[1] = "1"
	map1[2] = "2"
	map1[3] = "3"
	map1[4] = "4"
	map1[5] = "5"

	fmt.Println(map1)

	value, ok := map1[8]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println(ok)
	}

	//删除数据
	delete(map1, 3)
	fmt.Println(map1)

	fmt.Println(len(map1))

	// 遍历
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	keys := make([]int, 0, len(map1))
	fmt.Println(keys)

	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	// 排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, v := range keys {
		fmt.Println(v, map1[v])
	}
}
