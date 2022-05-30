package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num1 := rand.Int()
	fmt.Println(num1)

	for i := 0; i < 10; i++ {
		num := rand.Intn(10)
		fmt.Println(num)
	}

	// 设置时间种子
	rand.Seed(1000)
	num2 := rand.Intn(10)
	fmt.Println(num2)
}
