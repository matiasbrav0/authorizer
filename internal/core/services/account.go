package services

import (
	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
	"github.com/mbravovaisma/authorizer/pkg/constants"
	"github.com/mbravovaisma/authorizer/pkg/violations"
)

type account struct {
	repository ports.AuthorizerRepository
}

func NewAccount(repository ports.AuthorizerRepository) ports.Account {
	return &account{
		repository: repository,
	}
}

func (a *account) CreateAccount(account *domain.Account) (domain.Response, error) {
	if a.repository.AccountExist(constants.AccountID) {
		accountData := a.repository.GetAccountData(constants.AccountID)

		return domain.Response{
			Account:    accountData.Account,
			Violations: []string{violations.AccountAlreadyInitialized},
		}, nil
	}

	response := a.repository.AccountCreate(constants.AccountID, account)

	return response.AccountMovements[0], nil
}
