package transactionsrv

import (
	"time"

	"github.com/mbravovaisma/authorizer/internal/core/config"

	"github.com/mbravovaisma/authorizer/pkg/log"
	"go.uber.org/zap"

	v "github.com/mbravovaisma/authorizer/internal/core/constants"
	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
	"github.com/mbravovaisma/authorizer/pkg/constants"
)

type service struct {
	repository ports.AuthorizerRepository
}

func New(repository ports.AuthorizerRepository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) PerformTransaction(amount int64, merchant string, time time.Time) (domain.Movement, error) {
	/* Violates the account-not-initialized */
	if !s.repository.Exist(constants.AccountID) {
		return domain.Movement{
			Account:    nil,
			Violations: []string{v.AccountNotInitialized},
		}, nil
	}

	/* Create an empty array of violations */
	violations := make([]string, 0)

	/* Get account */
	account, err := s.repository.Get(constants.AccountID)
	if err != nil {
		log.Error("error getting account", zap.Error(err))
		return domain.Movement{}, err
	}

	/* Violates card-not-active */
	if !account.HasActiveCard() {
		violations = append(violations, v.CardNotActive)
	}

	/* Violates insufficient-limit */
	if !account.HasEnoughAmount(amount) {
		violations = append(violations, v.InsufficientLimit)
	}

	/* Violates the high-frequency-small-interval */
	if account.CanMakeATransaction(time) {
		violations = append(violations, v.HighFrequencySmallInterval)
	}

	/* Make a transaction */
	transaction := domain.NewTransaction(amount, merchant, time)

	/* Violates the doubled-transaction */
	if isDuplicatedTransaction(account, transaction) {
		violations = append(violations, v.DoubledTransaction)
	}

	/* If a violation occurred, don't execute transaction */
	if len(violations) > 0 {
		return domain.Movement{
			Account:    &account,
			Violations: violations,
		}, nil
	}

	/* Processing a transaction in happy path */
	execute(&account, transaction)

	err = s.saveAccountIntoRepository(account)
	if err != nil {
		log.Error("error saving account", zap.Error(err))
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

func isDuplicatedTransaction(account domain.Account, transaction domain.Transaction) bool {
	if account.Transactions == nil {
		return false
	}
	// [trx 1, trx2, trx3, trx4]
	for len(account.Transactions) > 0 {
		length := len(account.Transactions) - 1

		lastTransaction := account.Transactions[length]
		if transaction.Time.Before(lastTransaction.Time.Add(config.HighFrequencySmallIntervalTime)) {
			if transaction.Amount == lastTransaction.Amount && transaction.Merchant == lastTransaction.Merchant {
				return true
			}
		} else {
			return false
		}

		account.Transactions = account.Transactions[:length]
	}

	return false
}

func execute(account *domain.Account, transaction domain.Transaction) {
	/* Subtract amount from available limit */
	account.AvailableLimit -= transaction.Amount

	/* Generate a movement */
	movement := domain.Movement{
		Account:    account,
		Violations: []string{},
	}
	account.Movements = append(account.Movements, movement)

	/* Save transaction */
	account.Transactions = append(account.Transactions, transaction)

	/* Should be update last authorization time? */
	if account.ViolatesTheIntervalToPerformATransaction(transaction.Time) {
		account.Attempts = 0
		account.AuthorizationTime = transaction.Time
	}

	/* Add an attempt */
	account.Attempts += 1
}
