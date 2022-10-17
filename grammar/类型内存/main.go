package main

import "fmt"

func main() {
	var age int = 10
	// 地址
	fmt.Println(&age)

	// 定义指针变量
	var ptr *int = &age
	fmt.Println(*ptr)
	fmt.Println(ptr)
	fmt.Println(&ptr)
}
