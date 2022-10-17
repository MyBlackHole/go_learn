package main

import "fmt"

func main() {
	// 声明
	var age int
	// 赋值
	age = 18
	fmt.Printf("%d\n", age)

	// 声明 赋值
	var age2 int = 19
	fmt.Printf("%d\n", age2)

	//  // 赋值类型不匹配编译报错
	// var age3 int = 10.01
	// fmt.Printf("%d\n", age3)

	// 第一种使用方式
	var num int = 18
	fmt.Println(num)

	// 第二种, 使用默认值
	var num2 int
	fmt.Println(num2)

	// 第三种, 自动类型推断
	var num3 = 10
	fmt.Println(num3)

	// 第四种, 自动类型推断
	num4 := 10
	fmt.Println(num4)

	var num5, num6 int
	fmt.Println(num5, num6)
}
