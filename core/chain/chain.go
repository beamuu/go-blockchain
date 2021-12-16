package chain

import (
	"go-blockchain/core/block"
)
type Blockchain struct {
	blocks []*block.Block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := block.CreateBlock(data, prevBlock.Hash)
	// append new block to the chain
	chain.blocks = append(chain.blocks, newBlock)
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*block.Block{block.Genesis()}}
}

func (chain *Blockchain) GetBlock(index int) *block.Block {
	return chain.blocks[index]
}

