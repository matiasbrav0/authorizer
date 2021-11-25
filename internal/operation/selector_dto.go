package operation

import (
	"time"

	"github.com/mbravovaisma/authorizer/internal/core/domain"
)

type AccountsFields struct {
	ActiveCard     bool  `json:"active-card"`
	AvailableLimit int64 `json:"available-limit"`
}

type AccountOperation struct {
	Account AccountsFields `json:"account"`
}

type TransactionFields struct {
	Amount   int64     `json:"amount"`
	Merchant string    `json:"merchant"`
	Time     time.Time `json:"time"`
}

type TransactionOperation struct {
	Transaction TransactionFields `json:"transaction"`
}

type Response domain.Movement

func BuildResponse(model domain.Movement) Response {
	return Response(model)
}
