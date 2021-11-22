package domain

/* Account request */
type Account struct {
	Account accountInfo `json:"account"`
}

type accountInfo struct {
	ActiveCard     bool  `json:"active-card"`
	AvailableLimit int64 `json:"available-limit"`
}
