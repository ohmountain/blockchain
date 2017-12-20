package main

import (
	"fmt"
	"renshan/blockchain/proto"
)

func main() {
	chain := new(proto.Chain)

	chain.TheCreation()

	chain.AppendBlock(proto.CreateBlock([]byte("Hello World")))
	chain.AppendBlock(proto.CreateBlock([]byte("你好世界")))

	for _, b := range chain.Blocks {
		fmt.Printf("%s\n", b.Data)
	}
}
