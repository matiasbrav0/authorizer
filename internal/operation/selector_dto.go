package operation

import "github.com/mbravovaisma/authorizer/internal/core/domain"

type AccountOperation struct {
	Account domain.Account `json:"account"`
}

type TransactionOperation struct {
	Transaction domain.Transaction `json:"transaction"`
}

type Response domain.Movement

func BuildResponse(model domain.Movement) Response {
	return Response(model)
}
