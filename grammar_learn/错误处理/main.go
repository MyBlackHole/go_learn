package main

import "fmt"

func main() {
	// test_recover()

	func() {
		defer func() {
			fmt.Println("defer func")
			if err := recover(); err != nil {
				fmt.Println("recover success")
			}
		}()

		arr := []int{1, 2, 3}
		fmt.Println(arr[4])
		fmt.Println("after panic")

	}()
	fmt.Println("after recover")
	// arr := []int{1, 2, 3}
	// fmt.Println(arr[4])
	// fmt.Println("after panic")
}

func test_recover() {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
	fmt.Println("after panic")
}
