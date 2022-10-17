package main

import "fmt"

func init() {
	fmt.Printf("init 函数1\n")
}

func init() {
	fmt.Printf("init 函数2\n")
}

func main() {
	fmt.Print("Hello World\n")
}
