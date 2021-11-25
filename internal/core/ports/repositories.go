package ports

import "github.com/mbravovaisma/authorizer/internal/core/domain"

type AccountRepository interface {
	Get(id string) (domain.Account, error)
	Save(id string, account domain.Account) error
	Exist(id string) bool
}
