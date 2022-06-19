package blc

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	// 区块高度
	Height int64
	// 上一个区块连HASH
	PrevBlockHash []byte
	// 交易数据
	Data []byte
	// 时间戳
	Timestamp int64
	// 当前 Hash
	Hash []byte
	// Nonce
	Nonce int64
}

func (block *Block) SetHash() {
	heightBytes := IntToHex(block.Height)
	timestampStr := strconv.FormatInt(block.Timestamp, 2)
	timestampBytes := []byte(timestampStr)
	blockBytes := bytes.Join(
		[][]byte{
			heightBytes,
			block.PrevBlockHash,
			block.Data,
			timestampBytes,
			block.Hash,
		}, []byte{},
	)
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}

// 创建新区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{
		Height:        height,
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
		Nonce:         0,
	}

	// 工作量证明
	pow := NewProofOfWork(block)
	// 挖矿验证
	hash, nonce := pow.Run()
	fmt.Printf("%x\n", hash)
	block.Hash = hash[:]
	block.Nonce = nonce
	// block.SetHash()
	return block
}

func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, make([]byte, 32, 32))
}
