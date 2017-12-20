package proto

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"testing"
)

func TestCreateBlock(t *testing.T) {

	// 测试块Data、Hash、PrevBlockHash是否正确

	data := []byte("Hello World")

	chain := new(Chain)
	block := CreateBlock(data)

	chain.TheCreation()
	chain.AppendBlock(block)

	if bytes.Equal(block.Data, data) == false {
		t.Error("Expected true, got", false)
		t.SkipNow()
	}

	if bytes.Equal(chain.TheCreationBlock().Hash, block.PrevBlockHash) == false {
		t.Error("Expected true, got", false)
		t.SkipNow()
	}

	hash := sha256.Sum256(bytes.Join([][]byte{block.PrevBlockHash, block.Data, []byte(strconv.FormatInt(block.Timestamp, 10))}, []byte{}))
	if bytes.Equal(block.Hash, hash[:]) == false {
		t.Error("Expected true, got", false)
		t.SkipNow()
	}
}
