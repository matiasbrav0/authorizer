package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccount_HasActiveCard(t *testing.T) {
	// Setup
	account := NewAccount(true, 0)

	// Execute
	result := account.HasActiveCard()

	// Verify
	assert.Equal(t, result, true)
}

func TestAccount_HasEnoughAmount(t *testing.T) {
	// Setup
	account := NewAccount(true, 10)

	// Execute
	resultTrue := account.HasEnoughAmount(5)
	resultFalse := account.HasEnoughAmount(20)

	// Verify
	assert.Equal(t, resultTrue, true)
	assert.Equal(t, resultFalse, false)
}

func TestAccount_CanMakeATransactionSuccessfully(t *testing.T) {
	// Setup
	account := NewAccount(true, 0)
	account.AuthorizationTime = time.Now()
	account.Attempts = 2

	transaction := NewTransaction(0, "test_merchant", account.AuthorizationTime.Add(30*time.Second))

	// Execute
	result := account.CanMakeATransaction(transaction)

	// Verify
	assert.Equal(t, result, true)
}

func TestAccount_CanMakeATransactionExceedAttemptButValidTime(t *testing.T) {
	// Setup
	account := NewAccount(true, 0)
	account.AuthorizationTime = time.Now()
	account.Attempts = 4

	transaction := NewTransaction(
		0,
		"test_merchant",
		account.AuthorizationTime.Add(maxTimeToDuplicateATransaction+1*time.Second),
	)

	// Execute
	result := account.CanMakeATransaction(transaction)

	// Verify
	assert.Equal(t, result, true)
}

func TestAccount_CanMakeATransactionValidAttemptBufInvalidTime(t *testing.T) {
	// Setup
	account := NewAccount(true, 0)
	account.AuthorizationTime = time.Now()
	account.Attempts = 2

	transaction := NewTransaction(
		0,
		"test_merchant",
		account.AuthorizationTime.Add(maxTimeToDuplicateATransaction-1*time.Second),
	)

	// Execute
	result := account.CanMakeATransaction(transaction)

	// Verify
	assert.Equal(t, result, true)
}

func TestAccount_CanMakeATransactionFalse(t *testing.T) {
	// Setup
	account := NewAccount(true, 0)
	account.AuthorizationTime = time.Now()
	account.Attempts = 4

	transaction := NewTransaction(
		0,
		"test_merchant",
		account.AuthorizationTime.Add(maxTimeToDuplicateATransaction-1*time.Second),
	)

	// Execute
	result := account.CanMakeATransaction(transaction)

	// Verify
	assert.Equal(t, result, false)
}

func TestAccount_IsDuplicatedTransaction(t *testing.T) {
	// Setup
	transaction1 := NewTransaction(0, "test_merchant", time.Now())
	transaction2 := NewTransaction(0, "test_merchant", time.Now())

	account := NewAccount(true, 0)
	account.Transactions = append(account.Transactions, *transaction1)

	// Execute
	result := account.IsDuplicatedTransaction(transaction2)

	// Verify
	assert.Equal(t, result, true)
}

func TestAccount_IsDuplicatedTransactionFalse(t *testing.T) {
	// Setup
	transaction1 := NewTransaction(1, "test_merchant", time.Now())
	transaction2 := NewTransaction(0, "test_merchant", time.Now())

	account := NewAccount(true, 0)
	account.Transactions = append(account.Transactions, *transaction1)

	// Execute
	result := account.IsDuplicatedTransaction(transaction2)

	// Verify
	assert.Equal(t, result, false)
}

func TestAccount_ExecuteTransaction(t *testing.T) {
	// Setup
	availableLimit := int64(10)
	transactionAmount := int64(5)
	account := NewAccount(true, availableLimit)
	transaction1 := NewTransaction(transactionAmount, "test_merchant", time.Now())

	// Execute
	account.ExecuteTransaction(transaction1)

	// Verify
	assert.Equal(t, account.AvailableLimit, availableLimit-transactionAmount)
}
