package domain

/* Account response */
type Response struct {
	Account    Account  `json:"account"`
	Violations []string `json:"violations"`
}

/* Object to save in mapper */
type AccountData struct {
	AccountInfo      *Account
	AccountMovements []Response
	TransactionsInfo []Transaction
}
