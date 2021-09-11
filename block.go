package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Timestamp    int64
	Transactions []*Transaction
	PreBlockHash []byte
	Hash         []byte
	Nonce        int
}

func NewBlock(Transactions []*Transaction, preBlockHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Transactions: Transactions,
		PreBlockHash: preBlockHash,
		Hash:         []byte{},
	}
	pow := NewProofOfWOrk(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}
func NewFirstBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}
func Deserialize(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
