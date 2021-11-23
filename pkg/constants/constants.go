package constants

import "time"

const (
	/*
	 * If I had to manage multiple accounts this id would be the
	 * identifier of each account. As it a single account, the id is hardcode.
	 */
	AccountID = "0"

	/* It's the minimum time between 3 successfully transactions. */
	HighFrequencySmallIntervalTime = 2 * time.Minute

	/* Max transactions should be processed in HighFrequencySmallIntervalTime */
	MaxAttempts = uint8(3)
)
