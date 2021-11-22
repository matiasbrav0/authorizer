package ports

import "github.com/mbravovaisma/authorizer/internal/core/domain"

type Selector interface {
	OperationSelector([]byte) (domain.Response, error)
}
