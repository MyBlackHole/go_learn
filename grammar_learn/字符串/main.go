package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "golang 好的"

	fmt.Println(len(str))

	for _, v := range str {
		fmt.Printf("%c\n", v)
	}

	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i])
	}

	// 字符串转整数
	num1, _ := strconv.Atoi("123")
	fmt.Println(num1)

	// 整数转字符串
	s := strconv.Itoa(123)
	fmt.Println(s)

	// 统计一个字符串有几个指定的子串
	count := strings.Count("blackhole", "kh")
	fmt.Println(count)

	// 不区分大小写的字符串比较
	flag := strings.EqualFold("blackhole", "BLACKHOLE")
	fmt.Println(flag)

	// 区分大小写的字符串比较
	fmt.Println("black" == "BLACK")

	// 返回子串在字符串第一次出现的索引值
	index := strings.Index("gogago", "go")
	fmt.Println(index)

	slice1 := []byte{65, 66, 67, 68, 69}
	s3 := string(slice1)
	fmt.Println(s3)

	s4 := "abcdef"
	slice2 := []byte(s4)
	fmt.Println(slice2)
}
