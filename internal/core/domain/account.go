package domain

import (
	"time"

	"github.com/mbravovaisma/authorizer/internal/core/config"
)

type Account struct {
	ActiveCard        bool          `json:"active-card"`
	AvailableLimit    int64         `json:"available-limit"`
	Movements         []Movement    `json:"-"`
	Transactions      []Transaction `json:"-"`
	AuthorizationTime time.Time     `json:"-"`
	Attempts          uint8         `json:"-"`
}

func NewAccount(activeCard bool, availableLimit int64) Account {
	return Account{
		ActiveCard:     activeCard,
		AvailableLimit: availableLimit,
		Movements:      []Movement{},
		Transactions:   []Transaction{},
		Attempts:       0,
	}
}

func (account *Account) HasActiveCard() bool {
	return account.ActiveCard
}

func (account *Account) HasEnoughAmount(amount int64) bool {
	return account.AvailableLimit >= amount
}

func (account *Account) CanMakeATransaction(transactionTime time.Time) bool {
	return account.Attempts >= config.MaxAttempts && !account.ViolatesTheIntervalToPerformATransaction(transactionTime)
}

func (account *Account) ViolatesTheIntervalToPerformATransaction(transactionTime time.Time) bool {
	return transactionTime.After(account.AuthorizationTime.Add(config.HighFrequencySmallIntervalTime))
}
