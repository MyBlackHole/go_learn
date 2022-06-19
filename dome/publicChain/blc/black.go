package blc

import (
	"bytes"
	"crypto/sha256"
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
}

func (block *Block) SetHash() {
	heightBytes := IntToHex(block.Height)
	timestampStr := strconv.FormatInt(block.Timestamp, 2)
	timestampBytes := []byte(timestampStr)
	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timestampBytes, block.Hash}, []byte{})
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}

// 创建新区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := Block{
		Height:        height,
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
	}
	block.SetHash()
	return &block
}
