package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

func main() {
	var n int32
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	fmt.Println(n)
}
