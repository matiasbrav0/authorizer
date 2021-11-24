package ports

import (
	"time"

	"github.com/mbravovaisma/authorizer/internal/core/domain"
)

type AccountService interface {
	Create(activeCard bool, availableLimit int64) (domain.Movement, error)
}

type TransactionService interface {
	PerformTransaction(amount int64, merchant string, time time.Time) (domain.Movement, error)
}
