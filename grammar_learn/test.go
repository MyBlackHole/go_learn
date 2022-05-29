package main

import (
	"fmt"
	"github.com/axgle/mahonia"
)

//var count = 10

func main() {
	var a = 1
	var b = a
	enc := mahonia.NewEncoder("gbk")
	fmt.Println(nil)
	fmt.Printf(enc.ConvertString("a：%d, b：%d"), &a, &b)
}
