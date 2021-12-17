package core

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockchainIterator struct {
	CurrentHash []byte
	Database *badger.DB
}

func (chain *Blockchain) AddBlock(data string) {
	var lastHash []byte
	
	err := chain.Database.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte("lh"))
		err = item.Value(func(v []byte) error {
			lastHash = v
			return nil
		})
		return err
	})
	Handle(err)

	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {

		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err

	})
	Handle(err)
	
}

func InitBlockchain() *Blockchain {
	var lastHash []byte
	opts := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opts)

	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis block has been prooved and will be added the the chain")

			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)

			err = txn.Set([]byte("lh"), genesis.Hash)
			Handle(err)

			lastHash = genesis.Hash

			return err

		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			err = item.Value(func(v []byte) error {
				lastHash = v
				return nil
			})
			return err
		}

	})
	Handle(err)

	blockchain := &Blockchain{lastHash, db}
	return blockchain
}

func (chain *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{chain.LastHash, chain.Database}
	
}
func (iter *BlockchainIterator) Next() *Block {
	var block *Block
	var encodedBlock []byte

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err)
		// might cause some weird error because of copying function
		err = item.Value(func (v []byte) error{
			encodedBlock = v
			return nil
		})
		block = Deserialize(encodedBlock)
		return err
	})
	Handle(err)
	fmt.Printf("%s\n", block.Data)
	iter.CurrentHash = block.PrevHash

	return block
}

// func (iter *BlockchainIterator) GetDBLastHash() *Block {
// 	var encodedBlock []byte
// 	var block *Block
// 	err := iter.Database.View(func(txn *badger.Txn) error {
// 		item, err := txn.Get(iter.CurrentHash)
// 		Handle(err)
// 		// might cause some weird error because of copying function
// 		_, err = item.ValueCopy(encodedBlock)
// 		block = Deserialize(encodedBlock)
// 		return err
// 	})
// 	Handle(err)
// 	return block
// }

