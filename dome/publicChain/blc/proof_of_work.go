package blc

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

const TARGETBIT = 19

type ProofOfWork struct {
	// 需要验证的区块
	Block *Block
	// 大数据存储
	target *big.Int
}

func (proofOfWork *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			proofOfWork.Block.PrevBlockHash,
			proofOfWork.Block.Data,
			IntToHex(proofOfWork.Block.Timestamp),
			IntToHex(int64(TARGETBIT)),
			IntToHex(int64(nonce)),
			IntToHex(proofOfWork.Block.Height),
		},
		[]byte{},
	)
	return data
}

func (proofOfWork *ProofOfWork) Run() (hash [32]byte, nonce int64) {
	var hashInt *big.Int
	hashInt = new(big.Int)
	for {
		// 将 Block 的属性拼接成字节数组
		dataBytes := proofOfWork.prepareData(nonce)
		// 生成 hash
		hash = sha256.Sum256(dataBytes)

		// fmt.Printf("%x\n", hash)

		// 将 hash 存储到 hashInt
		hashInt.SetBytes(hash[:])
		// 判断 hash 有效性， 满足条件跳出循环
		if proofOfWork.target.Cmp(hashInt) == 1 {
			break
		}
		nonce++
	}
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
