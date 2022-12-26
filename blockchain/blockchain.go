package blockchain

import (
	"fmt"

	"github.com/diyliv/blockchain/block"
)

type BlockChain struct {
	Blocks []block.Block `json:"blocks"`
}

func NewBlockChain(bl ...block.Block) *BlockChain {
	blocks := make([]block.Block, 0)
	blocks = append(blocks, bl...)

	return &BlockChain{
		Blocks: blocks,
	}
}

func (b *BlockChain) Print() {
	for _, val := range b.Blocks {
		fmt.Printf("Hash: %x\nPrevHash: %x\nTransactions: %v\nCreatedAt: %v\n\n", val.Hash, val.PrevHash, val.Transaction, val.CreatedAt)
	}
}
