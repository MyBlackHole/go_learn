package main

import (
	"fmt"
	"publicChain/blc"
)

func main() {
	genesisBlockChain := blc.CreateBlockChainWithGenesisBlock()
	fmt.Println(genesisBlockChain)
}
