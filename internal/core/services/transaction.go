package services

import (
	"time"

	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
	"github.com/mbravovaisma/authorizer/pkg/constants"
	"github.com/mbravovaisma/authorizer/pkg/violations"
)

type transaction struct {
	repository ports.AuthorizerRepository
}

func NewTransaction(repository ports.AuthorizerRepository) ports.Transaction {
	return &transaction{
		repository: repository,
	}
}

func (t *transaction) PerformTransaction(transaction *domain.Transaction) (domain.Response, error) {
	/* Violates the account-not-initialized */
	if !t.repository.AccountExist(constants.AccountID) {
		return domain.Response{
			Account:    nil,
			Violations: []string{violations.AccountNotInitialized},
		}, nil
	}

	/* Get account */
	authorizerData := t.repository.GetAccountData(constants.AccountID)

	/* Violates card-not-active */
	if !t.isActiveCards(authorizerData.Account) {
		return domain.Response{
			Account:    authorizerData.Account,
			Violations: []string{violations.CardNotActive},
		}, nil
	}

	/* Violates insufficient-limit */
	if !t.hasValidAmount(authorizerData.Account, transaction) {
		return domain.Response{
			Account:    authorizerData.Account,
			Violations: []string{violations.InsufficientLimit},
		}, nil
	}

	/* Violates the high-frequency-small-interval */
	if t.isHighFrequencySmallInterval(authorizerData, transaction) {
		return domain.Response{
			Account:    authorizerData.Account,
			Violations: []string{violations.HighFrequencySmallInterval},
		}, nil
	}

	/* Processing a transaction in happy path */
	r := t.executeTransaction(authorizerData, transaction)

	return domain.Response{
		Account:    r.Account,
		Violations: []string{},
	}, nil
}

func (t *transaction) isHighFrequencySmallInterval(accountData *domain.AuthorizerData, transaction *domain.Transaction) bool {
	return accountData.Attempts >= constants.MaxAttempts &&
		!t.isHighFrequencySmallIntervalTime(accountData.AuthorizationTime, transaction.Time)
}

func (t *transaction) isHighFrequencySmallIntervalTime(authorizationTime time.Time, transactionTime time.Time) bool {
	return transactionTime.After(authorizationTime.Add(constants.HighFrequencySmallIntervalTime))
}

// hasValidAmount return true if transaction has valid amount to perform a transaction, otherwise return false
func (t *transaction) hasValidAmount(account *domain.Account, transaction *domain.Transaction) bool {
	return account.AvailableLimit >= transaction.Amount
}

// isActiveCards return true if account has an active card or false if has an inactive card
func (t *transaction) isActiveCards(account *domain.Account) bool {
	return account.ActiveCard
}

func (t *transaction) executeTransaction(accountData *domain.AuthorizerData, transaction *domain.Transaction) *domain.AuthorizerData {
	/* Subtract amount from available limit */
	accountData.Account.AvailableLimit -= transaction.Amount

	/* Generate a movement */
	movement := domain.Response{
		Account:    accountData.Account,
		Violations: []string{},
	}

	/* Should be update last authorization time? */
	if t.isHighFrequencySmallIntervalTime(accountData.AuthorizationTime, transaction.Time) {
		accountData.Attempts = 0
		accountData.AuthorizationTime = transaction.Time
	}

	/* Add an attempt */
	accountData.Attempts += 1

	/* If execute is successful sync movement with current data */
	return t.repository.UpdateAccountData(
		constants.AccountID,
		accountData.Account,
		transaction,
		&movement,
		accountData.AuthorizationTime,
		accountData.Attempts,
	)
}
