package main

// /*
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L./lib/ -lgm3000.1.0
// #include "gm3000.1.0.h"
// */
// import "C"
// import (
// 	"fmt"
// 	"strings"
// 	"unsafe"
// )

// // go build -ldflags="-r ./" 1.go

// func get_ukey() (string) {
// 	var ulRslt C.uint
// 	szDevName := make([]byte, 256)
// 	var ulNameLen C.uint = 256
// 	ulRslt = C.SKF_EnumDev(1, (*C.char)(unsafe.Pointer(&szDevName[0])), &ulNameLen)
// 	a := string(szDevName)
// 	fmt.Println(ulRslt)
// 	fmt.Println(a)
// 	fmt.Println(len(a))
// 	str := strings.Replace(a, "\x00", "", -1)
// 	return str
// }
