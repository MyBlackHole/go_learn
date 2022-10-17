package main

import "fmt"

func main() {
	if true {
		defer func() {
			fmt.Println("defer ")
		}()
	}

	// 匿名函数
	defer func() {
		fmt.Println("defer func")
	}()

	defer fun1("Black Hole")

	arr := []int{1, 2, 3}
	fmt.Println(arr[0])
	test()
	test2()
	fmt.Println("来了")
}

func fun1(s string) {
	fmt.Println(s)
}

func test() {
	defer fun1("ok")
	i := 100

	defer func(i1 int) {
		fmt.Println(i1)
	}(i)

	defer fmt.Println(i)
	i += 100

	defer func() {
		fmt.Println(i)
	}()
	return
}

func test2() {
	defer func() {
		fmt.Printf("o\n")
		if err := recover(); err != nil {
			fmt.Printf("%s\n", err)
		}
	}()
	i := 1
	i1 := 0
	i3 := i / i1
	fmt.Println("1 / 0")
	fmt.Println(i3)
}
