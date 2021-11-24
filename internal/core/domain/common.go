package domain

type Movement struct {
	Account    *Account `json:"account"`
	Violations []string `json:"violations"`
}
