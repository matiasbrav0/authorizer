package services

import (
	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
)

type transaction struct {
	repository ports.AuthorizerRepository
}

func NewTransaction(repository ports.AuthorizerRepository) ports.Transaction {
	return &transaction{
		repository: repository,
	}
}

func (t transaction) PerformTransaction(transaction domain.Transaction) (domain.Response, error) {
	panic("implement me")
}
