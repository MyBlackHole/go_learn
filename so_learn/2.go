package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lgm3000.1.0
#include "gm3000.1.0.h"
*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

// go build -ldflags="-r ./" 1.go

func main() {
	var ulRslt C.uint
	//var szDevName C.char
	szDevName := make([]byte, 256)
	var ulNameLen C.uint = 256
	ulRslt = C.SKF_EnumDev(1, (*C.char)(unsafe.Pointer(&szDevName[0])), &ulNameLen)
	a := string(szDevName)
	fmt.Println(ulRslt)
	fmt.Println(a)
	fmt.Println(len(a))
	str := strings.Replace(a, "\x00", "", -1)
	fmt.Println(str)
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
