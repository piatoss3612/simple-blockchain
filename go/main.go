package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	hash     string
	prevHash string
	data     string
	height   int
}

func NewBlock(prevHash string, data string, height int) Block {
	block := Block{
		hash:     calculateHash(prevHash, data, height),
		prevHash: prevHash,
		data:     data,
		height:   height,
	}
	return block
}

func calculateHash(prevHash string, data string, height int) string {
	hash := sha256.New()
	_, _ = hash.Write([]byte(fmt.Sprintf("%s%d%s", prevHash, height, data)))
	return hex.EncodeToString(hash.Sum(nil))
}

type Blockchain struct {
	blocks []Block
}

func NewBlockchain() Blockchain {
	return Blockchain{
		blocks: make([]Block, 0),
	}
}

func (bc *Blockchain) AddBlock(data string) {
	prevHash := bc.prevHash()
	height := len(bc.blocks) + 1
	block := NewBlock(prevHash, data, height)
	bc.blocks = append(bc.blocks, block)
}

func (bc *Blockchain) prevHash() string {
	if len(bc.blocks) == 0 {
		return ""
	}

	return bc.blocks[len(bc.blocks)-1].hash
}

func (bc Blockchain) GetBlocks() []Block {
	return bc.blocks
}

func main() {
	blockchain := NewBlockchain()

	blockchain.AddBlock("First block")
	blockchain.AddBlock("Second block")
	blockchain.AddBlock("Third block")

	blocks := blockchain.GetBlocks()

	for _, block := range blocks {
		fmt.Printf("Prev. hash: %s\n", block.prevHash)
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %s\n", block.hash)
		fmt.Println("==================================")
	}
}
