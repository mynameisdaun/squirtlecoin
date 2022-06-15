package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", b.getLastHash()}
	newBlock.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(newBlock.data+newBlock.prevHash)))
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) listBlocks() {
	for idx, block := range b.blocks {
		fmt.Printf("index: %d, Data : %s, Curr Hash: %s, Prev Hash: %s\n", idx, block.data, block.hash, block.prevHash)
	}
}

func main() {
	chain := blockchain{}
	chain.addBlock("first")
	chain.addBlock("second")
	chain.listBlocks()
}
