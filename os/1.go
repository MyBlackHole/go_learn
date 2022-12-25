package main

import (
	"fmt"
	"os"
)

func main() {
	var HOME string
	HOME = os.Getenv("HOME")
	fmt.Println(HOME)
	// fmt.Sprintf("%s/%s", HOME, path)
	err := os.RemoveAll("test")
	if err != nil {
		fmt.Println(err.Error())
	}
}
