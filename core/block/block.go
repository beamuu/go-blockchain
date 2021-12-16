package block

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	data := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(data)
	b.Hash = hash[:] 
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	// Calculate block hash
	block.DeriveHash()
	// return as a pointer
	return block
}
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}