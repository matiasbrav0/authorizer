package domain

import "time"

type Transaction struct {
	Transaction transactionInfo `json:"transaction"`
}

type transactionInfo struct {
	Merchant string    `json:"merchant"`
	Amount   int64     `json:"amount"`
	Time     time.Time `json:"time"`
}
