package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	func1()
}

func func1() {
	data, err := ioutil.ReadFile("语法_learn\\text.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	} else {
		fmt.Println("Contents of file:", string(data))
	}
}
