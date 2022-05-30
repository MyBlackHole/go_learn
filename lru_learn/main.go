package main

import (
	"fmt"

	"github.com/hashicorp/golang-lru"
)

func main() {
	l, _ := lru.New(128)
	for i := 0; i < 256; i++ {
		l.Add(i, nil)
	}

	if l.Len() != 128 {
		panic(fmt.Sprintf("%d", l.Len()))
	}
}
