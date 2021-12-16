package core

type Blockchain struct {
	Blocks []*Block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	// append new block to the chain
	chain.Blocks = append(chain.Blocks, newBlock)
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func (chain *Blockchain) GetBlock(index int) *Block {
	return chain.Blocks[index]
}

