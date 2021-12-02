package dependencies

import (
	"github.com/mbravovaisma/authorizer/internal/core/ports"
	"github.com/mbravovaisma/authorizer/internal/core/services/accountsrv"
	"github.com/mbravovaisma/authorizer/internal/core/services/allowlistsrv"
	"github.com/mbravovaisma/authorizer/internal/core/services/transactionsrv"
	"github.com/mbravovaisma/authorizer/internal/operation"
	"github.com/mbravovaisma/authorizer/internal/repositories/accountrepo"
)

type definitions struct {
	// Repositories
	Repository ports.AccountRepository

	// Services
	AccountService     ports.AccountService
	TransactionService ports.TransactionService
	AllowListService   ports.AllowListService

	// Selector (Driver adapter)
	Selector *operation.Selector
}

func NewDependencies() *definitions {
	memKvs := accountrepo.New()
	accountService := accountsrv.New(memKvs)
	transactionService := transactionsrv.New(memKvs)
	allowListService := allowlistsrv.New(memKvs)
	selector := operation.NewSelector(accountService, transactionService, allowListService)

	return &definitions{
		Repository:         memKvs,
		AccountService:     accountService,
		TransactionService: transactionService,
		AllowListService:   allowListService,
		Selector:           selector,
	}
}
