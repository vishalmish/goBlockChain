package main

import "time"

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce 		  int
}

/**
 * Create a new Block. This can then be added through AddBlock
 */
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}



/**
 * For a brand new Blockchain, we need to have 1 genesis block to start addding
 * more blocks.
 */
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}