package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/diyliv/blockchain/block"
	"github.com/diyliv/blockchain/transactions"
)

func main() {
	firstBlock := block.Block{
		Transaction: []transactions.Transaction{
			{Sender: "A", Recipient: "B", Amount: 1337.1337},
		},
	}
	m, err := json.Marshal(firstBlock)
	if err != nil {
		panic(err)
	}

	firstBlockHash := sha256.Sum256(m)
	firstBlock.Hash = fmt.Sprintf("%s", firstBlockHash)
	firstBlock.PrevHash = firstBlock.Hash

	firstBlock.NewBlock().Print()

	secondBlock := block.Block{
		Transaction: []transactions.Transaction{
			{Sender: "C", Recipient: "D", Amount: 0.1337},
		},
	}

	secM, err := json.Marshal(secondBlock)
	if err != nil {
		panic(err)
	}

	secondBlockHash := sha256.Sum256(secM)
	secondBlock.Hash = fmt.Sprintf("%s", secondBlockHash)
	secondBlock.PrevHash = firstBlock.Hash
	secondBlock.NewBlock().Print()
}
