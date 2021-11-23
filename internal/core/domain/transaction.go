package domain

import "time"

type TransactionRequest struct {
	Transaction Transaction `json:"transaction"`
}

type Transaction struct {
	Merchant string    `json:"merchant"`
	Amount   int64     `json:"amount"`
	Time     time.Time `json:"time"`
}
