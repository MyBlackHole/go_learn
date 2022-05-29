package main

import "fmt"


func main() {
	if true {
		defer func() {
			fmt.Println("defer ")
		}()
	}
	defer func() {
		fmt.Println("defer func")
	}()
	arr := []int{1, 2, 3}
	fmt.Println(arr[0])
}
