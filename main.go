package main

import (
	"fmt"
	"renshan/blockchain/proto"
)

func main() {
	chain := new(proto.Chain)

	// 创世
	chain.TheCreation()

	chain.AppendBlock(proto.CreateBlock([]byte("Hello World")))
	chain.AppendBlock(proto.CreateBlock([]byte("你好世界")))

	fmt.Println("The Creation Block' Hash: ", chain.TheCreationBlock().GetHashString())
}
