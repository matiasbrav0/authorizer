package config

import "time"

const (
	/* It's the minimum time between 3 successfully transactions. */
	HighFrequencySmallIntervalTime = 2 * time.Minute

	/* Max transactions should be processed in HighFrequencySmallIntervalTime */
	MaxAttempts = uint8(3)
)
