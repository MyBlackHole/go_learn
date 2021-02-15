package main

import "fmt"

type Pint struct {
	x int
	y int
}

func main() {
	var a Pint = Pint{1, 1}
	b := Pint{1, 1}
	fmt.Println(a, b)
}
