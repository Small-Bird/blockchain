package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"

const blockBucket = "blocks"

type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

func (bc *BlockChain) AddBlock(Transactions []*Transaction) {
	var lastHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	newBlock := NewBlock(Transactions, lastHash)
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
func (bc *BlockChain) Iterator() *BlockChainIterator {
	bci := &BlockChainIterator{bc.tip, bc.db}
	return bci
}
func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}
func CreateBlockChain(address string) *BlockChain {
	if dbExists() {
		fmt.Println("存在数据库文件")
		os.Exit(1)
	}
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		cbtx := NewCoinbaseTx(address, "创币交易")
		genesis := NewFirstBlock(cbtx)
		b, err := tx.CreateBucket([]byte(blockBucket))
		if err != nil {
			log.Panic(err)
		}
		err = b.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"), genesis.Hash)
		tip = genesis.Hash
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc := BlockChain{tip, db}
	return &bc
}
func NewBlockChain() *BlockChain {
	if dbExists() == false {
		fmt.Println("不存在数据库文件，请先创建数据库文件")
		os.Exit(1)
	}
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		tip = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc := BlockChain{tip, db}
	return &bc
}
func (bc *BlockChain) FindUnspectedTransactions(address string) []Transaction {
	var unspentTxs []Transaction
	spentTxs := make(map[string][]int)
	bci := bc.Iterator()
	for {
		block := bci.Next()
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)
		OutPuts:
			for index, out := range tx.Vout {
				if spentTxs[txID] != nil {
					for _, spendOut := range spentTxs[string(index)] {
						if spendOut == index {
							continue OutPuts
						}
					}
				}
				if out.CanBeUnlockedWith(address) {
					unspentTxs = append(unspentTxs, *tx)
				}
			}
			if tx.IsCoinBase() == false {
				for _, in := range tx.Vin {
					if in.CanUnlockOutputWith(address) {
						inTxId := hex.EncodeToString(in.Txid)
						spentTxs[inTxId] = append(spentTxs[inTxId], in.Vout)
					}
				}
			}
		}
		if len(block.PreBlockHash) == 0 {
			break
		}
	}
	return unspentTxs
}
func (tx *Transaction) IsCoinBase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}
