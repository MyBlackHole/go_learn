package main

import (
	"fmt"
	"os"
)

func main() {
	var HOME string
	HOME = os.Getenv("HOME")
	fmt.Println(HOME)
}
