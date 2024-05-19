package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestPseudoRandomGenerator(t *testing.T) {
	tests := []struct {
		name      string
		start     int
		end       int
		n         int
		wantError bool
	}{
		{"ValidRange", 1, 10, 5, false},
		{"StartGreaterThanEnd", 10, 1, 5, true},
		{"NegativeCount", 1, 10, -5, true},
		{"ZeroCount", 1, 10, 0, true},
		{"SingleNumberRange", 5, 5, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pseudoRandomGenerator(tt.start, tt.end, tt.n, " ")
			if (err != nil) != tt.wantError {
				t.Errorf("pseudoRandomGenerator() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !tt.wantError {
				// Split the result into individual numbers
				numbers := strings.Split(got, " ")
				if len(numbers) != tt.n {
					t.Errorf("expected %d numbers, got %d", tt.n, len(numbers))
				}

				for _, numStr := range numbers {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						t.Errorf("failed to convert %s to int: %v", numStr, err)
					}
					if num < tt.start || num > tt.end {
						t.Errorf("number %d out of range [%d, %d]", num, tt.start, tt.end)
					}
				}
			}
		})
	}
}
