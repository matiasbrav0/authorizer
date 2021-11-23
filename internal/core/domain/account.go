package domain

type AccountRequest struct {
	Account Account `json:"account"`
}

type Account struct {
	ActiveCard     bool  `json:"active-card"`
	AvailableLimit int64 `json:"available-limit"`
}
