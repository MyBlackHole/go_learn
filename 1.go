package main

import "fmt"

var iii int

func test() (i int, j int) {
	i = 1
	j = 2
	return
}

func main() {
	i1, j1 := test()
	fmt.Println(i1, j1)
	fmt.Println(iii)
}
