package blc

type BlockChain struct {
	// 有序区块
	Blocks []*Block
}

// 创建带有创世区块的区块连
func CreateBlockChainWithGenesisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis Data")
	return &BlockChain{[]*Block{genesisBlock}}
}
