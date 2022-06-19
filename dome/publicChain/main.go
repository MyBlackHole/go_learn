package main

import (
	"fmt"
	"publicChain/blc"
)

func main() {
	// blockChain := blc.CreateBlockChainWithGenesisBlock()

	// blockChain.Add("send 100RMB", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	// blockChain.Add("send 400RMB", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	// blockChain.Add("send 800RMB", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	// fmt.Println(blockChain)

	block := blc.NewBlock("test", 1, make([]byte, 32, 32))
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	proofOfWork := blc.NewProofOfWork(block)
	fmt.Printf("%v\n", proofOfWork.IsValid())
}