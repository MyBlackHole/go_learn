package main

import "fmt"

func main() {
	result := func(num1, num2 int) int {
		return num1 + num2
	}(10, 20)

	fmt.Println(result)

	sub := func(num1 int, num2 int) int {
		return num1 - num2
	}

	result1 := sub(80, 100)
	fmt.Println(result1)
}
