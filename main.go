package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
)


func main() {

	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Vishal")
	bc.AddBlock("Send 2 more BTC to Vishal")

	for _, block := range bc.blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

