package main

import "fmt"

func main() {
	var n1 int = +10
	fmt.Println(n1)

	var n2 int = 4 + 7
	fmt.Println(n2)

	var s1 string = "abc" + "def"
	fmt.Println(s1)

	fmt.Println(10 / 3)
	// 3
	fmt.Println(10.0 / 3)
	// 3.3333333333333335

	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	fmt.Println(!false)
	fmt.Println(!true)
}
