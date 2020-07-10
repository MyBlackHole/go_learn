package main

import "fmt"

//func main() {
//	arr := []int{1, 2, 3}
//	fmt.Println(arr[4])
//}

func main() {
	fmt.Println("before panic")
	panic("crash")
	fmt.Println("after panic")
}
