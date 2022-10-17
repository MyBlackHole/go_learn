package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// 可变参数当作切片处理
func test(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func test1(num int) {
	num = 30
	fmt.Println(num)
}

func test2(num *int) {
	*num = 30
	fmt.Println(*num)
}

func test02(a int) {
	fmt.Println(a)
}

func test05(i int) (ret int) {
	ret = i + i
	return
}

func main() {
	var num1 int = 10
	var num2 int = 20

	var sum int

	sum += num1
	sum += num2
	fmt.Println(sum)

	fmt.Println(add(40, 39))

	test()
	test(3)
	test(3, 3)

	var num int = 10
	test1(num)
	fmt.Println(num)

	test2(&num)
	fmt.Println(num)

	a := test
	fmt.Printf("%T, %T\n", a, test)

	aa := test02
	aa(1)
}
