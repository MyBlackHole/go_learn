package main

import "fmt"

func main() {
	var f func()
	f = fun1
	f()

	arr1 := fun2()
	fmt.Printf("%T-- %p-- %v\n", arr1, &arr1, arr1)

	arr3 := fun3()
	fmt.Printf("%T-- %p-- %v\n", arr3, &arr3, arr3)
}

func fun3() *[4]int {
	return &[4]int{1, 2, 3, 4}
}

func fun2() [4]int {
	return [4]int{1, 2, 3, 4}
}

func fun1() {
	fmt.Println("ok")
}
