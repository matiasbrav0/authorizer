package services

import (
	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
)

type account struct {
	repository ports.AuthorizerRepository
}

func NewAccount(repository ports.AuthorizerRepository) ports.Account {
	return &account{
		repository: repository,
	}
}

func (a *account) CreateAccount(account domain.Account) (domain.AccountResponse, error) {
	panic("implement me")
}
