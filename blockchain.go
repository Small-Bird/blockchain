package main

import (
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"

const blockBucket = "blocks"

type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		if b == nil {
			genesiz := NewFirstBlock()
			b, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic(err)
			}
			err = b.Put(genesiz.Hash, genesiz.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = b.Put([]byte("l"), genesiz.Hash)
			tip = genesiz.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc := BlockChain{tip, db}
	return &bc
}
func (bc *BlockChain) AddBlock(data string) {
	var lastHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	newBlock := NewBlock(data, lastHash)
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("1"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.tip = newBlock.Hash
		return nil
	})
}
