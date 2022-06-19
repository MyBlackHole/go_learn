package blc

type ProofOfWork struct {
	Block *Block
}

func (proofOfWork *ProofOfWork) Run() (hash []byte, nonce int64) {
	return
}

// 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	return &ProofOfWork{Block: block}
}
