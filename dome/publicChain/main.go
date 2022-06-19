package main

import (
	"fmt"
	"publicChain/blc"
)

func main() {
	blockChain := blc.CreateBlockChainWithGenesisBlock()
	fmt.Println(blockChain)

	blockChain.Add("send 100RMB", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	fmt.Println(blockChain)

	blockChain.Add("send 400RMB", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	fmt.Println(blockChain)

	blockChain.Add("send 800RMB", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	fmt.Println(blockChain)
}
