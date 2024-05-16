package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello, World!", "!dlroW ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"日本語", "語本日"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := reverseString(tt.input)
			if result != tt.expected {
				t.Errorf("reverseString(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
