package main

import (
	"testing"
)

func TestConvertUnixTimestamp(t *testing.T) {
	format := "2006-01-02 15:04:05.000000000"
	tests := []struct {
		timestamp int64
		expected  string
	}{
		{1627670400005, "2021-07-30 20:40:00.005000000"},
		{1627670400000005, "2021-07-30 20:40:00.000005000"},
		{1627670400000000005, "2021-07-30 20:40:00.000000005"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result, err := convertUnixTimestamp(tt.timestamp, format)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("convertUnixTimestamp(%d, %q) = %q; want %q", tt.timestamp, format, result, tt.expected)
			}
		})
	}
}

func TestConvertUnixTimestampInvalid(t *testing.T) {
	format := "2006-01-02 15:04:05.999999999"
	_, err := convertUnixTimestamp(1627670, format)
	if err == nil {
		t.Error("expected error for invalid timestamp, got nil")
	}
}
