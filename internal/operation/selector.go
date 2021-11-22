package operation

import "github.com/mbravovaisma/authorizer/internal/core/ports"

type selector struct {
	accountService     ports.Account
	transactionService ports.Transaction
}

func NewSelector(accountService ports.Account, transactionService ports.Transaction) ports.Selector {
	return &selector{
		accountService:     accountService,
		transactionService: transactionService,
	}
}

func (s selector) OperationSelector(operation []byte) error {
	panic("implement me")
}
