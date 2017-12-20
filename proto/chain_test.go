package proto

import "testing"
import "bytes"

func TestTheCreation(t *testing.T) {
	chain := new(Chain)

	if chain.TheCreationBlock() != nil {
		t.Error("Expected true, got", false)
	}

	chain.TheCreation()
	chain.TheCreation()
	chain.TheCreation()

	if len(chain.Blocks) != 1 {
		t.Error("Expected 1, got", len(chain.Blocks))
		t.SkipNow()
	}

	if bytes.Equal(chain.TheCreationBlock().Data, []byte("The Creation Block")) == false {
		t.Error("Expected true, got", false)
		t.SkipNow()
	}

	if bytes.Equal(chain.Blocks[0].Data, []byte("The Creation Block")) == false {
		t.Error("Expected true, got", false)
		t.SkipNow()
	}

	if bytes.Equal(chain.Blocks[0].PrevBlockHash, nil) == false {
		t.Error("Expected true, got", false)
		t.SkipNow()
	}
}

func TestGetLastBlock(t *testing.T) {

	chain := new(Chain)
	chain.TheCreation()

	if chain.Blocks[0] != chain.GetLastBLock() {
		t.Error("Expected true, got", false)
		t.SkipNow()
	}
}

func TestAppendBlock(t *testing.T) {
	chain := new(Chain)
	chain.TheCreation()

	// 测试链的长度是否正确
	// 测试最新添加的是否是最后的
	// 测试PrevBockHash是否正确

	chain.AppendBlock(CreateBlock([]byte("测试 AppendBlock")))

	if len(chain.Blocks) != 2 {
		t.Error("Expected true, got", false)
	}

	chain.AppendBlock(CreateBlock([]byte("Test AppendBlock")))

	if len(chain.Blocks) != 3 {
		t.Error("Expected true, got", false)
	}

	for i, b := range chain.Blocks {
		if i == 0 {
			continue
		}

		if bytes.Equal(chain.Blocks[i-1].Hash, b.PrevBlockHash) == false {
			t.Error("Expected true, got", false)
		}
	}
}
