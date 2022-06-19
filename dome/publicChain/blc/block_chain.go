package blc

type BlockChain struct {
	// 有序区块
	Blocks []*Block
}

func (blockChain *BlockChain) Add(data string, heigh int64, preHash []byte) {
	// 创建新区块
	newBlock := NewBlock(data, heigh, preHash)
	// 往链里添加区块
	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

// 创建带有创世区块的区块连
func CreateBlockChainWithGenesisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis Data")
	return &BlockChain{[]*Block{genesisBlock}}
}
