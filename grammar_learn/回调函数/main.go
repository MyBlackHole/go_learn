package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	oper(2, 4, add)
}

func oper(a, b int, fun func(a, b int) int) int {
	// fmt.Println(a, b, fun)
	res := fun(a, b)
	fmt.Println(res)
	var e error
	defer func() { // 释放锁
		t(e)
	}()
	e = fmt.Errorf("eeeeee")

	return res
}

func t(e error) {
	if e != nil {
		fmt.Println(e.Error(), "---")
	} else {
		fmt.Println("ok")
	}
}

func add(a, b int) int {
	return a + b
}
