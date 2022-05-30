package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	oper(2, 4, add)
}

func oper(a, b int, fun func(a, b int) int) int {
	fmt.Println(a, b, fun)
	res := fun(a, b)
	fmt.Println(res)
	return res
}

func add(a, b int) int {
	return a + b
}
