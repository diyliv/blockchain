package transactions

type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
}

func NewTransaction(sender, recipient string, amount float64) *Transaction {
	return &Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
}
