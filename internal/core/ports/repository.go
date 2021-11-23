package ports

import "github.com/mbravovaisma/authorizer/internal/core/domain"

type AuthorizerRepository interface {
	AccountCreate(string, *domain.Account) *domain.AccountData
	AccountExist(string) bool
	UpdateAccountData(string, *domain.Account, *domain.Transaction, *domain.Response) *domain.AccountData
	GetAccountData(string) *domain.AccountData
}
