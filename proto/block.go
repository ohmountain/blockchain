package proto

import (
	"encoding/hex"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	Hash          []byte
	PrevBlockHash []byte
	Nonce         int
	PoF           *ProofOfWork
}

func (b *Block) hashing() {

	pof := NewProofOfWork(b)
	nonce, hash := pof.Run()

	b.Hash = hash[:]
	b.Nonce = nonce
}

func (b *Block) SetPrevBlockHash(prevBlockHash []byte) {
	b.PrevBlockHash = prevBlockHash

	defer b.hashing()
}

func (b *Block) GetHashString() string {
	return hex.EncodeToString(b.Hash)
}

func (b *Block) Validate() bool {
	return b.PoF.Validate()
}

func CreateBlock(data []byte) *Block {
	block := new(Block)

	block.Data = data
	block.Timestamp = time.Now().Unix()

	return block
}
