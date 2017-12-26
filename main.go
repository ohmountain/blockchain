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

	for _, b := range chain.Blocks {
		fmt.Println("")
		fmt.Println("The Creation Block' Hash: ", b.GetHashString())
		fmt.Println("The Creation Block' Nonce: ", b.Nonce)
		fmt.Println("The Creation Block' Validate: ", b.Validate())
	}
}
