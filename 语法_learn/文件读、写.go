package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	func2()
}

func func2() {
	f, err := os.Create("语法_learn\\text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Black Hole")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "ok")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func func1() {
	data, err := ioutil.ReadFile("语法_learn\\text.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	} else {
		fmt.Println("Contents of file:", string(data))
	}
}
