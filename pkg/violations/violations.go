package violations

const (
	// Once created, the account should not be updated or recreated
	AccountAlreadyInitialized = "account-already-initialized"

	// No transaction should be accepted without a properly initialized account
	AccountNotInitialized = "account-not-initialized"

	// No transaction should be accepted when the card is not active
	CardNotActive = "card-not-active"

	// The transaction amount should not exceed the available limit
	InsufficientLimit = "insufficient-limit"

	// There should be no more than 3 transactions within a 2 minutes interval
	HighFrequencySmallInterval = "high-frequency-small-interval"

	// There should be no more than 1 similar transaction (same amount and merchant) within a 2 minutes interval
	DoubledTransaction = "doubled-transaction"

	// Unexpected error
	UnexpectedViolation = "unexpected-violation"
)
