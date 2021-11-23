package domain

import "time"

type Transaction struct {
	Transaction TransactionInfo `json:"transaction"`
}

type TransactionInfo struct {
	Merchant string    `json:"merchant"`
	Amount   int64     `json:"amount"`
	Time     time.Time `json:"time"`
}
