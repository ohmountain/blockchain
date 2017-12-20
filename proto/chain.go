package proto

type Chain struct {
	Blocks []*Block
}

func (c *Chain) TheCreation() {
	if len(c.Blocks) == 0 {
		block := CreateBlock([]byte("The Creation Block"))
		block.SetPrevBlockHash([]byte{})
		c.Blocks = append(c.Blocks, block)
	}
}

func (c *Chain) TheCreationBlock() *Block {
	if len(c.Blocks) == 0 {
		return nil
	}

	return c.Blocks[0]
}

func (c *Chain) GetLastBLock() *Block {
	if len(c.Blocks) == 0 {
		return nil
	}

	return c.Blocks[len(c.Blocks)-1]
}

func (c *Chain) AppendBlock(b *Block) {
	if len(c.Blocks) > 0 {
		b.SetPrevBlockHash(c.GetLastBLock().Hash)
		c.Blocks = append(c.Blocks, b)
	}
}
