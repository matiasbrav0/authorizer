package domain

import (
	"time"
)

// Account configs
const (
	// It's the minimum time between 3 successfully transactions
	highFrequencySmallIntervalTimeToPerformATransaction = 2 * time.Minute

	// Max transactions should be processed in HighFrequencySmallIntervalTime
	maxAttempts = uint8(3)

	// Max time to accept transactions with same merchant and amount
	maxTimeToDuplicateATransaction = 2 * time.Minute
)

type Account struct {
	ActiveCard        bool          `json:"active-card"`
	AvailableLimit    int64         `json:"available-limit"`
	AllowListed       bool          `json:"allow-listed"`
	Movements         []Movement    `json:"-"`
	Transactions      []Transaction `json:"-"`
	AuthorizationTime time.Time     `json:"-"`
	Attempts          uint8         `json:"-"`
}

func NewAccount(activeCard bool, availableLimit int64) *Account {
	return &Account{
		ActiveCard:     activeCard,
		AvailableLimit: availableLimit,
		AllowListed:    false,
		Movements:      []Movement{},
		Transactions:   []Transaction{},
		Attempts:       0,
	}
}

func (a *Account) HasActiveCard() bool {
	return a.ActiveCard
}

func (a *Account) HasEnoughAmount(amount int64) bool {
	return a.AvailableLimit >= amount
}

func (a *Account) CanMakeATransaction(transaction *Transaction) bool {
	if a.isAllowListed() {
		return true
	}

	return a.Attempts < maxAttempts || a.notViolatesTheIntervalToPerformATransaction(transaction.Time)
}

func (a *Account) SetAllowList(value bool) *Account {
	a.AllowListed = value
	return a
}

func (a *Account) IsDuplicatedTransaction(transaction *Transaction) bool {
	if a.isAllowListed() {
		return false
	}

	if a.Transactions == nil {
		return false
	}

	index := len(a.Transactions) - 1
	for index > -1 {
		lastTransaction := a.Transactions[index]

		// Only going to analyze the transactions that were made before of maxTimeToDuplicateATransaction (2 minutes)
		if transaction.Time.Before(lastTransaction.Time.Add(maxTimeToDuplicateATransaction)) {
			if transaction.Amount == lastTransaction.Amount && transaction.Merchant == lastTransaction.Merchant {
				return true
			}
		} else {
			return false
		}

		index -= 1
	}

	return false
}

func (a *Account) ExecuteTransaction(transaction *Transaction) {
	// Subtract amount from available limit
	a.AvailableLimit = a.AvailableLimit - transaction.Amount

	// Generate a movement
	movement := Movement{
		Account:    a,
		Violations: []string{},
	}
	a.Movements = append(a.Movements, movement)

	// Save transaction
	a.Transactions = append(a.Transactions, *transaction)

	// Should be update last authorization time?
	if a.notViolatesTheIntervalToPerformATransaction(transaction.Time) {
		a.Attempts = 0
		a.AuthorizationTime = transaction.Time
	}

	// Add an attempt
	a.Attempts += 1
}

// ----- Private functions ----- //

func (a *Account) notViolatesTheIntervalToPerformATransaction(transactionTime time.Time) bool {
	return transactionTime.After(a.AuthorizationTime.Add(highFrequencySmallIntervalTimeToPerformATransaction))
}

func (a *Account) isAllowListed() bool {
	return a.AllowListed
}
