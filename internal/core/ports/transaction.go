package ports

import "github.com/mbravovaisma/authorizer/internal/core/domain"

type Transaction interface {
	PerformTransaction(*domain.Transaction) (domain.Response, error)
}
