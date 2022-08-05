package main

import "fmt"

func main() {
	a := true
	T(a)
	fmt.Println(a)
	T(a)
	fmt.Println(a)
}

func T(a bool) {
	if a {
		fmt.Println(a)
		fmt.Println("o")
		a = false
	}
}
