package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L./so/ -ltest
#include "test.h"
*/
import "C"
import (
	"fmt"
)

// go build -ldflags="-r ./so/" 1.go 

func main() {
	a := C.add(1, 3)
	fmt.Println(a)
	//ptr, err := plugin.Open("/home/black/go/src/go_learn/so_learn/libtest.so")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//Add, _ := ptr.Lookup("add")
	//sum := Add.(func(int, int) int)(5, 4)
	//fmt.Println("Add结果：", sum)
	//
	//Sub, _ := ptr.Lookup("Subtract")
	//sub := Sub.(func(int, int) int)(9, 8)
	//fmt.Println("Sub结果：", sub)
}

