package proto

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	Hash          []byte
	PrevBlockHash []byte
}

func (b *Block) hashing() {
	hash := sha256.Sum256(bytes.Join([][]byte{b.PrevBlockHash, b.Data, []byte(strconv.FormatInt(b.Timestamp, 10))}, []byte{}))
	b.Hash = hash[:]
}

func (b *Block) SetPrevBlockHash(prevBlockHash []byte) {
	b.PrevBlockHash = prevBlockHash

	defer b.hashing()
}

func CreateBlock(data []byte) *Block {
	block := new(Block)

	block.Data = data
	block.Timestamp = time.Now().Unix()

	return block
}
