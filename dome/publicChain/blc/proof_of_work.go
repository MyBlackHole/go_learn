package blc

import "math/big"

const TARGETBIT = 16

type ProofOfWork struct {
	// 需要验证的区块
	Block *Block
	// 大数据存储
	target *big.Int
}

func (proofOfWork *ProofOfWork) Run() (hash []byte, nonce int64) {
	return
}

// 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	// 创建为 1 的 target
	target := big.NewInt(1)
	// 左移 256 - targetBit
	target = target.Lsh(target, 256-TARGETBIT)
	return &ProofOfWork{Block: block, target: target}
}
