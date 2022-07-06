package main

import (
	"fmt"
	"unsafe"
)

func main() {
	buckets := make([]unsafe.Pointer, 10)
	for i, v := range buckets {
		fmt.Println(int32(i), v)
	}
}
