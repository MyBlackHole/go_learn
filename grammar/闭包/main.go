package main

import "fmt"

func getSum() func(int) int {
	var sum int = 10
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

func main() {
	f := getSum()
	fmt.Println(f(10))
	fmt.Println(f(20))
}
