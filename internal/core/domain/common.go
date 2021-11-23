package domain

import "time"

type Response struct {
	Account    *Account `json:"account"`
	Violations []string `json:"violations"`
}

/* Object to save in mapper */
type AuthorizerData struct {
	Account           *Account
	AccountMovements  []Response
	TransactionsInfo  []Transaction
	AuthorizationTime time.Time
	Attempts          uint8
}
