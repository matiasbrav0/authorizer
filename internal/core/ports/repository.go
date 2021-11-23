package ports

import (
	"time"

	"github.com/mbravovaisma/authorizer/internal/core/domain"
)

type AuthorizerRepository interface {
	AccountCreate(string, *domain.Account) *domain.AuthorizerData
	AccountExist(string) bool
	UpdateAccountData(string, *domain.Account, *domain.Transaction, *domain.Response, time.Time, uint8) *domain.AuthorizerData
	GetAccountData(string) *domain.AuthorizerData
}
