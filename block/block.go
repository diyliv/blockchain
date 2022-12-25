package block

import (
	"fmt"
	"time"

	"github.com/diyliv/blockchain/transactions"
)

type Block struct {
	Hash        string                     `json:"block_hash"`
	PrevHash    string                     `json:"previous_hash"`
	Transaction []transactions.Transaction `json:"transactions"`
	CreatedAt   time.Time                  `json:"created_at"`
}

func (b *Block) NewBlock() *Block {
	return &Block{
		Hash:        b.Hash,
		PrevHash:    b.PrevHash,
		Transaction: b.Transaction,
		CreatedAt:   time.Now().Local(),
	}
}

func (b *Block) Print() {
	fmt.Printf("Hash: %x\nPrevHash: %x\nTransactions: %v\nCreatedAt: %v\n", b.Hash, b.PrevHash, b.Transaction, b.CreatedAt)
}
