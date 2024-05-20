package main

import (
	"fmt"
	"time"
)

func convertUnixTimestamp(timestamp int64, format string) (string, error) {
	var t time.Time

	// Determine the precision based on the length of the timestamp
	switch {
	case timestamp >= 1e18:
		// Nanoseconds
		t = time.Unix(0, timestamp)
	case timestamp >= 1e15:
		// Microseconds
		t = time.UnixMicro(timestamp)
	case timestamp >= 1e12:
		// Milliseconds
		t = time.UnixMilli(timestamp)
	case timestamp >= 1e9:
		// Seconds
		t = time.Unix(timestamp, 0)
	default:
		return "", fmt.Errorf("invalid timestamp: %d", timestamp)
	}

	return t.Format(format), nil
}
