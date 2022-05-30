package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	num1 := rand.Int()
	fmt.Println(num1)
}
