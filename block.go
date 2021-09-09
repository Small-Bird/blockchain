package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	PreBlockHash []byte
	Hash         []byte
	Nonce        int
}

func (b *Block) SetHash() {
	timestap := IntToHex(b.Timestamp)
	headers := bytes.Join([][]byte{timestap, b.Data, b.PreBlockHash}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Data:         []byte(data),
		PreBlockHash: preBlockHash,
		Hash:         []byte{},
	}
	block.SetHash()
	pow := NewProofOfWOrk(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}
func NewFirstBlock() *Block {
	return NewBlock("第一个区块", []byte{})
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
