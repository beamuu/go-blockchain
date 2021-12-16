package core

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func (b *Block) DeriveHash() {
	data := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(data)
	b.Hash = hash[:] 
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	// Calculate block hash
	pow := NewProof(block)
	nonce, hash := pow.Run()
	
	block.Hash = hash[:]
	block.Nonce = nonce
	
	return block
}
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}