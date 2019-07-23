package main


type Blockchain struct {
	blocks []*Block
}

/**
 * Add a new block into the block chain.
 */
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}


/**
 * Create a new Blockchain with a NewGenesisBlock.
 */
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
