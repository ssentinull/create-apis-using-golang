package utils

import "time"

func ParseDuration(in string, defaultDuration time.Duration) time.Duration {
	dur, err := time.ParseDuration(in)
	if err != nil {
		return defaultDuration
	}

	return dur
}
