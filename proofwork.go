package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 20

var maxNonce = math.MaxInt64

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWOrk(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}
func (pow *ProofOfWork) prepareData(nonce int) []byte {

	data := bytes.Join([][]byte{
		pow.block.PreBlockHash,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce)),
	}, []byte{},
	)
	var vinData []byte
	var voutData []byte

	bytesBuffer1 := bytes.NewBuffer([]byte{})
	bytesBuffer2 := bytes.NewBuffer([]byte{})
	bytesBuffer3 := bytes.NewBuffer([]byte{})

	for _, transaction := range pow.block.Transactions {
		for _, vin := range (*transaction).Vin {
			binary.Write(bytesBuffer1, binary.BigEndian, vin.ScriptSig)
			binary.Write(bytesBuffer2, binary.BigEndian, vin.Vout)

			vinData = bytes.Join([][]byte{data, vin.Txid, bytesBuffer1.Bytes(), bytesBuffer2.Bytes()}, []byte{})
		}
		for _, vout := range (*transaction).Vout {
			binary.Write(bytesBuffer3, binary.BigEndian, vout.Value)

			voutData = bytes.Join([][]byte{vinData, bytesBuffer3.Bytes()}, []byte{})
		}
	}
	return voutData
}
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce += 1
		}
	}

	fmt.Println("\n")
	return nonce, hash[:]

}
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
