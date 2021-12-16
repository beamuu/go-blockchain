package main

import (
	"fmt"
	"go-blockchain/core/chain"
)

func main() {
	Chain := chain.InitBlockchain()
	
	// Interact with blockchain here

	Chain.AddBlock("block1")
	Chain.AddBlock("block2")

	fmt.Println(string(Chain.GetBlock(0).Data))

}
