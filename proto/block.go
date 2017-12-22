package proto

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	Hash          []byte
	PrevBlockHash []byte
	Nonce         int
	PoW           *ProofOfWork
}

func (b *Block) hashing() {

	pow := NewProofOfWork(b)
	nonce, hash := pow.Run()

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
	return b.PoW.Validate()
}

func (b *Block) Serialize() ([]byte, error) {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(b); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

func Deserialize(b []byte) (*Block, error) {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(b))

	if err := decoder.Decode(&block); err != nil {
		return nil, err
	}

	return &block, nil
}

func CreateBlock(data []byte) *Block {
	block := new(Block)

	block.Data = data
	block.Timestamp = time.Now().Unix()

	return block
}
