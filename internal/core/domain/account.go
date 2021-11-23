package domain

/* Account request */
type Account struct {
	Account AccountInfo `json:"account"`
}

type AccountInfo struct {
	ActiveCard     bool  `json:"active-card"`
	AvailableLimit int64 `json:"available-limit"`
}
