package domain

import "time"

type Transaction struct {
	Amount   int64     `json:"amount"`
	Merchant string    `json:"merchant"`
	Time     time.Time `json:"time"`
}

func NewTransaction(amount int64, merchant string, time time.Time) Transaction {
	return Transaction{
		Amount:   amount,
		Merchant: merchant,
		Time:     time,
	}
}
