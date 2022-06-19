package main

import (
	"fmt"
	"unsafe"
)

type W struct {
	b int32
	c int64
}

func main() {
	var w *W = new(W)
	//这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b, w.c)

	//现在我们通过指针运算给b变量赋值为10
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int64)(b)) = 10
	//此时结果就变成了10，0
	fmt.Println(w.b, w.c)
}
