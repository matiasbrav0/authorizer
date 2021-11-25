package transactionsrv

import (
	"time"

	"github.com/mbravovaisma/authorizer/pkg/log"

	"github.com/mbravovaisma/authorizer/internal/core/constants"
	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
)

type service struct {
	repository ports.AccountRepository
}

func New(repository ports.AccountRepository) ports.TransactionService {
	return &service{
		repository: repository,
	}
}

func (s *service) PerformTransaction(amount int64, merchant string, time time.Time) (domain.Movement, error) {
	// Violates the account-not-initialized
	if !s.repository.Exist(constants.AccountID) {
		return domain.Movement{
			Account:    nil,
			Violations: []string{constants.AccountNotInitialized},
		}, nil
	}

	// Create an empty array of violations
	violations := make([]string, 0)

	// Get account
	account, err := s.repository.Get(constants.AccountID)
	if err != nil {
		log.Error("error getting account", log.ErrorField(err))
		return domain.Movement{}, err
	}

	// Violates card-not-active
	if !account.HasActiveCard() {
		violations = append(violations, constants.CardNotActive)
	}

	// Violates insufficient-limit
	if !account.HasEnoughAmount(amount) {
		violations = append(violations, constants.InsufficientLimit)
	}

	// Violates the high-frequency-small-interval
	if !account.CanMakeATransaction(time) {
		violations = append(violations, constants.HighFrequencySmallInterval)
	}

	// Make a transaction
	transaction := domain.NewTransaction(amount, merchant, time)

	// Violates the doubled-transaction
	if account.IsDuplicatedTransaction(transaction) {
		violations = append(violations, constants.DoubledTransaction)
	}

	// If a violation occurred, don't execute transaction
	if len(violations) > 0 {
		return domain.Movement{
			Account:    &account,
			Violations: violations,
		}, nil
	}

	// Processing a transaction in happy path
	account.ExecuteTransaction(transaction)

	err = s.saveAccountIntoRepository(account)
	if err != nil {
		log.Error("error saving account", log.ErrorField(err))
		return domain.Movement{}, err
	}

	return domain.Movement{
		Account:    &account,
		Violations: violations,
	}, nil
}

func (s *service) saveAccountIntoRepository(account domain.Account) error {
	return s.repository.Save(constants.AccountID, account)
}
