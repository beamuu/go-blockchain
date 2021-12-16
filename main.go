package main

import (
	"fmt"
	"go-blockchain/core"
	"strconv"
)

func main() {
	chain := core.InitBlockchain()
	
	// Interact eith blockchain here

	chain.AddBlock("block1")
	chain.AddBlock("block2")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := core.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
