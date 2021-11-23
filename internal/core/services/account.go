package services

import (
	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
	"github.com/mbravovaisma/authorizer/pkg/constants"
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

	}
	a.repository.AccountSave(constants.AccountID, account)
	return domain.Response{}, nil
}
