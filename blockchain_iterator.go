package main

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (i *BlockChainIterator) Next() *Block {
	var block *Block
	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		encodedBlock := b.Get(i.currentHash)
		block = Deserialize(encodedBlock)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	i.currentHash = block.PreBlockHash
	return block
}
