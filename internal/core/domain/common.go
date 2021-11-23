package domain

/* Account response */
type Response struct {
	Account    data     `json:"account"`
	Violations []string `json:"violations"`
}

type data struct {
	ActiveCard     bool `json:"active-card"`
	AvailableLimit int  `json:"available-limit"`
}

/* Object to save in mapper */
type AccountData struct {
	AccountInfo      *Account
	AccountMovements []Response
	TransactionsInfo []Transaction
}
