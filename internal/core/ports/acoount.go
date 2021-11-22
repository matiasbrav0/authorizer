package ports

import "github.com/mbravovaisma/authorizer/internal/core/domain"

type Account interface {
	CreateAccount(*domain.Account) (domain.Response, error)
}
