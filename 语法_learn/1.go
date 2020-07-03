package main

import "fmt"

func main() {
	func7()
}
func func7() {
	//逻辑运算符
	var a bool = true
	var b bool = false

	if a && b {
		fmt.Printf("line 1 - condition is true\n")
	}

	if a || b {
		fmt.Printf("line 2 - condition is true\n")
	}
	a = false
	b = true

	if a && b {
		fmt.Printf("line 3 - condition is true\n")
	} else {
		fmt.Printf("line 3 - condition is true\n")
	}

}
func func6() {
	//关系运算符
	var a int = 21
	var b int = 10
	if a == b {
		fmt.Printf("line 1 - a = b\n")
	} else {
		fmt.Printf("line 1 - a not= b\n")
	}

	if a < b {
		fmt.Printf("line 2 - a < b\n")
	} else {
		fmt.Printf("line 2 - a not< b\n")
	}

	if a > b {
		fmt.Printf("line 3 - a > b\n")
	} else {
		fmt.Printf("line 3 - a not> b\n")
	}

	a = 5
	b = 20

	if a <= b {
		fmt.Printf("line 2 - a < b\n")
	} else {
		fmt.Printf("line 2 - a not< b\n")
	}

	if a >= b {
		fmt.Printf("line 3 - a > b\n")
	} else {
		fmt.Printf("line 3 - a not> b\n")
	}
}
func func5() {
	//算数运算符
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf(" Line 1 value of c is %d\n", c)

	c = a - b
	fmt.Printf(" Line 2 value of c is %d\n", c)

	c = a * b
	fmt.Printf(" Line 3 value of c is %d\n", c)

	c = a / b
	fmt.Printf(" Line 4 value of c is %d\n", c)

	c = a % b
	fmt.Printf(" Line 5 value of c is %d\n", c)

	a++
	fmt.Printf(" Line 6 value of a is %d\n", a)

	a--
	fmt.Printf(" Line 6 value of a is %d\n", a)
}
func func4() {
	//const 关键字
	const LENGTH = 10
	const WIDTH = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area:%d", area)
}

func func3() {
	//变量声明
	var a, b, c = 3, 4, "foo"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
}

func func2() {
	//动态类型声明
	var x float64 = 20

	y := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
}

func func1() {
	//静态类型声明
	fmt.Println("Hello World")
}
