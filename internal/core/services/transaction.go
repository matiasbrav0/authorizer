package services

import "github.com/mbravovaisma/authorizer/internal/core/ports"

type transaction struct {
	repository ports.AuthorizerRepository
}

func NewTransaction(repository ports.AuthorizerRepository) ports.Transaction {
	return &transaction{
		repository: repository,
	}
}
