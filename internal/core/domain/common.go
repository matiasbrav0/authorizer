package domain

/* Account response */
type Response struct {
	Account    AccountInfo `json:"account"`
	Violations []string    `json:"violations"`
}

/* Object to save in mapper */
type AccountData struct {
	AccountInfo      *AccountInfo
	AccountMovements []Response
	TransactionsInfo []Transaction
}
