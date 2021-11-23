package services

import (
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
	account := t.repository.GetAccountData(constants.AccountID)

	/* Violates card-not-active */
	if !t.isActiveCards(account.AccountInfo) {
		return domain.Response{
			Account:    account.AccountInfo,
			Violations: []string{violations.CardNotActive},
		}, nil
	}

	/* Violates insufficient-limit */
	if !t.hasValidAmount(account.AccountInfo, transaction.Transaction) {
		return domain.Response{
			Account:    account.AccountInfo,
			Violations: []string{violations.InsufficientLimit},
		}, nil
	}

	/* Processing a transaction successfully */
	r := t.executeTransaction(transaction)

	return r, nil
}

// hasValidAmount return true if transaction has valid amount to perform a transaction, otherwise return false
func (t *transaction) hasValidAmount(account *domain.AccountInfo, transaction domain.TransactionInfo) bool {
	return account.AvailableLimit >= transaction.Amount
}

// isActiveCards return true if account has an active card or false if has an inactive card
func (t *transaction) isActiveCards(account *domain.AccountInfo) bool {
	return account.ActiveCard
}

func (t *transaction) executeTransaction(transaction *domain.Transaction) domain.Response {
	accountData := t.repository.GetAccountData(constants.AccountID)

	accountData.AccountInfo.AvailableLimit -= transaction.Transaction.Amount

	movement := domain.Response{
		Account:    accountData.AccountInfo,
		Violations: []string{},
	}

	t.repository.UpdateAccountData(constants.AccountID, accountData.AccountInfo, transaction, &movement)

	return movement
}
