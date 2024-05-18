package main

import (
	"testing"
)

func TestStringToHex(t *testing.T) {
	tests := []struct {
		input    string
		spaced   bool
		expected string
	}{
		{"hello world", false, "68656c6c6f20776f726c64"},
		{"hello world", true, "68 65 6c 6c 6f 20 77 6f 72 6c 64"},
		{"Hello, 世界", false, "48656c6c6f2c20e4b896e7958c"},
		{"Hello, 世界", true, "48 65 6c 6c 6f 2c 20 e4 b8 96 e7 95 8c"},
		{"", false, ""},
		{"", true, ""},
		{"abc", false, "616263"},
		{"abc", true, "61 62 63"},
		{"Go is awesome!", false, "476f20697320617765736f6d6521"},
		{"Go is awesome!", true, "47 6f 20 69 73 20 61 77 65 73 6f 6d 65 21"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := stringToHex(tt.input, tt.spaced)
			if result != tt.expected {
				t.Errorf("stringToHex(%q, %v) = %q; want %q", tt.input, tt.spaced, result, tt.expected)
			}
		})
	}
}

func TestHexToString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"68656c6c6f20776f726c64", "hello world"},
		{"48656c6c6f2c20e4b896e7958c", "Hello, 世界"},
		{"48 65 6c 6c 6f 2c 20 e4 b8 96 e7 95 8c", "Hello, 世界"},
		{"", ""},
		{"616263", "abc"},
		{"476f20697320617765736f6d6521", "Go is awesome!"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := hexToString(tt.input)
			if err != nil {
				t.Fatalf("hexToString(%q) returned an error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("hexToString(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
