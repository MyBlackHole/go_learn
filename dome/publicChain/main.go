package main

import (
	"fmt"
	"publicChain/blc"
)

func main() {
	prevBlockHash := make([]byte, 32, 32)
	block := blc.NewBlock("Genenis Block", 1, prevBlockHash)
	fmt.Println(block)
}
